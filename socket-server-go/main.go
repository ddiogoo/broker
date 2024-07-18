package main

import (
	"github.com/ddiogoo/broker/tree/master/socket-server-go/env"
	"github.com/ddiogoo/broker/tree/master/socket-server-go/server"
)

// Method responsible for initializing our application.
func main() {
	env.Config()
	err := server.Start()
	if err != nil {
		panic(err.Error())
	}
}
