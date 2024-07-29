package main

import (
	"os"
	"sync"

	"github.com/ddiogoo/broker/tree/master/socket-server-go/mq"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	n, err := mq.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer n.Close()

	go func() {
		defer wg.Done()
		err := n.Publish("hello", os.Args[1])
		if err != nil {
			panic(err.Error())
		}
	}()
	wg.Wait()
}
