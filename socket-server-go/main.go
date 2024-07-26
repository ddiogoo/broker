package main

import (
	"github.com/ddiogoo/broker/tree/master/socket-server-go/env"
	"github.com/ddiogoo/broker/tree/master/socket-server-go/mq"
	"github.com/ddiogoo/broker/tree/master/socket-server-go/server"
)

// main is the entry point of the application.
func main() {
	env.Config()
	_, err := mq.ConnectToNats()
	if err != nil {
		panic(err.Error())
	}

	err = server.Start()
	if err != nil {
		panic(err.Error())
	}
}
