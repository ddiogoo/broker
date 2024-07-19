package server

import (
	"errors"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	// Message about Error generated when the server try to read messages.
	ErrReadMessageWebSocket = "error on reading messages"
	// Message about Error generated when the client try to connect.
	ErrUpgradeWebSocketProtocol = "error on upgrade http connection"
	// Error generated when the application try to start the server.
	ErrStartWebSocketServer = errors.New("error on starting websocket server")
)
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var addr = flag.String("addr", "localhost:3000", "http service address")

// Start WebSocket Server or return an error.
func Start() error {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		return ErrStartWebSocketServer
	}
	return nil
}

// Responsible to upgrade Http connections and read messages.
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Fatalln(ErrReadMessageWebSocket)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
