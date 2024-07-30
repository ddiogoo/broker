package client

import (
	"log"
	"net/http"
	"time"

	"github.com/ddiogoo/broker/tree/master/socket-server-go/mq"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var (
	newline = []byte{'\n'}
)

// Upgrader is used to upgrade an HTTP connection to a WebSocket connection.
var clientUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Client represents a single WebSocket connection.
type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	username string
}

// fromHubToClientConnection pumps messages from the hub to the WebSocket connection.
func (c *Client) fromHubToClientConnection(ch chan string) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-ch:
			if !ok {
				log.Println("channel closed")
				return
			}
			log.Println("Received message:", message)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("NextWriter error:", err)
				return
			}
			if _, err := w.Write([]byte(message)); err != nil {
				log.Println("Write error:", err)
				return
			}
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				log.Println("Close writer error:", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				return
			}
		}
	}
}

// ServeWs handles WebSocket requests from the peer.
func ServeWs(n *mq.Nats, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := clientUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "anonymous"
	}
	c := &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
		username: username}
	c.hub.register <- c
	log.Println("Client registered:", username)

	ch := make(chan string)
	subscription := make(chan struct{})

	go func() {
		err := n.Subscribe("hello", ch)
		if err != nil {
			log.Fatalln("Error on subscription to 'hello':", err)
		}
		close(subscription)
	}()
	<-subscription
	log.Println("Subscribed to 'hello'")

	go c.fromHubToClientConnection(ch)
}
