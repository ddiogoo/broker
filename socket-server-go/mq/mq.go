package mq

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Config() {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, _ := nc.SubscribeSync("foo")
	_, _ = sub.NextMsg(nats.DefaultOptions.Timeout)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, _ = nc.ChanSubscribe("foo", ch)
	msg := <-ch
	log.Println(string(msg.Data))

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, _ = nc.Request("help", []byte("help me"), 10*time.Millisecond)
	log.Println(string(msg.Data))

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}