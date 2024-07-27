package mq

import (
	"errors"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var (
	// errConnectToNats error on try to connect to nats server.
	errConnectToNats = errors.New("error on try connection to nats server")
)

// Nats struct has Conn property to manager the nats.
type Nats struct {
	Conn *nats.Conn
}

// Publish send a msg to a subj.
func (n *Nats) Publish(subj string, msg string) {
	err := n.Conn.Publish(subj, []byte(msg))
	if err != nil {
		log.Println("error on publish a message on " + subj)
	}
}

// Subscriber receive messages from subj and print on terminal.
func (n *Nats) Subscribe(subj string) {
	n.Conn.Subscribe(subj, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}

// Connect try to connect to a Nats Server.
func Connect() (*Nats, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, errConnectToNats
	}
	return &Nats{Conn: nc}, nil
}
