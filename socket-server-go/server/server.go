package server

import (
	"flag"
	"net/http"

	"github.com/ddiogoo/broker/tree/master/socket-server-go/client"
)

// addr is a pointer to a string that represents the address of the server.
var addr = flag.String("addr", ":8080", "http service address")

// Start is responsible for starting the server.
func Start() error {
	flag.Parse()
	hub := client.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		return err
	}
	return nil
}
