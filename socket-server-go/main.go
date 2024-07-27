package main

import (
	"log"

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

	ch := make(chan string)
	subscription := make(chan struct{})
	go func() {
		err := n.Subscribe("hello", ch)
		if err != nil {
			log.Fatalln("error on subscription hello")
		}
		close(subscription)
	}()
	<-subscription

	go func() {
		for v := range ch {
			log.Println("Message received: " + v)
		}
	}()
	n.Publish("hello", "DIOGO!")

	err = server.Start()
	if err != nil {
		panic(err.Error())
	}
}
