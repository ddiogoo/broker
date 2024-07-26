package mq

import (
	"errors"

	"github.com/nats-io/nats.go"
)

var (
	errConnectToNats = errors.New("error on try connection to nats server")
)

// ConnectToNats try to connect to a Nats Server.
func ConnectToNats() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, errConnectToNats
	}
	return nc, nil
}

// func config() {
// 	// Connect to a server
// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	// Simple Publisher
// 	nc.Publish("foo", []byte("Hello World"))

// 	// Simple Async Subscriber
// 	nc.Subscribe("foo", func(m *nats.Msg) {
// 		fmt.Printf("Received a message: %s\n", string(m.Data))
// 	})

// 	// Responding to a request message
// 	nc.Subscribe("request", func(m *nats.Msg) {
// 		m.Respond([]byte("answer is 42"))
// 	})

// 	// Simple Sync Subscriber
// 	sub, err := nc.SubscribeSync("foo")
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	_, _ = sub.NextMsg(time.Second * 10)

// 	// Channel Subscriber
// 	ch := make(chan *nats.Msg, 64)
// 	sub, err = nc.ChanSubscribe("foo", ch)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	msg := <-ch
// 	log.Println(string(msg.Data))

// 	// Unsubscribe
// 	sub.Unsubscribe()

// 	// Drain
// 	sub.Drain()

// 	// Requests
// 	msg, err = nc.Request("help", []byte("help me"), 10*time.Millisecond)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	log.Println(string(msg.Data))

// 	// Replies
// 	nc.Subscribe("help", func(m *nats.Msg) {
// 		nc.Publish(m.Reply, []byte("I can help!"))
// 	})

// 	// Drain connection (Preferred for responders)
// 	// Close() not needed if this is called.
// 	nc.Drain()

// 	// Close connection
// 	nc.Close()
// }
