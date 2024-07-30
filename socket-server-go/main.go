package main

import (
	"github.com/ddiogoo/broker/tree/master/socket-server-go/env"
	"github.com/ddiogoo/broker/tree/master/socket-server-go/mq"
	"github.com/ddiogoo/broker/tree/master/socket-server-go/server"
)

// main is the entry point of the application.
func main() {
	env.Config()
	n, err := mq.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer n.Close()
	err = server.Start(n)
	if err != nil {
		panic(err.Error())
	}
}
