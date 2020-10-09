package main

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

const ()

func setupWebserver() {

	box := packr.NewBox("./ui/public")
	http.Handle("/", http.FileServer(box))
	_ = http.ListenAndServe(":8080", nil)

	/*
	   // Initialize WebSocket handler + server
	   mux := http.NewServeMux()
	       mux.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
	           func() {
	               for {
	                   fmt.Println("hi")
	               }
	           }
	       }
	   // messageType initializes some type of message
	   message := messageType{}
	   if err := websocket.JSON.Receive(conn, &message); err != nil {
	       // handle error
	   }
	   // send message
	   if err := websocket.JSON.Send(conn, message); err != nil {
	       // handle error
	   }
	*/
}
