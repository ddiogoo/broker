package mq

import (
	"errors"

	"github.com/nats-io/nats.go"
)

var (
	// errConnectToNats error on try to connect to nats server.
	errConnectToNats = errors.New("error on try connection to nats server")
)

// Nats struct has Conn property to manager the nats.
type Nats struct {
	conn *nats.Conn
}

// Close closes the connection.
func (n *Nats) Close() {
	n.conn.Close()
}

// Publish send a msg to a subj.
func (n *Nats) Publish(subj string, msg string) error {
	err := n.conn.Publish(subj, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

// Subscriber receive messages from subj and print on terminal.
func (n *Nats) Subscribe(subj string, ch chan (string)) error {
	_, err := n.conn.Subscribe(subj, func(m *nats.Msg) {
		ch <- string(m.Data)
	})
	return err
}

// Connect try to connect to a Nats Server.
func Connect() (*Nats, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, errConnectToNats
	}
	return &Nats{conn: nc}, nil
}
