package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rosso-ai/conlai/web"
)

func main() {
	hub := web.NewHub()
	go hub.Run()

	http.HandleFunc("/ws/{clientID}", func(w http.ResponseWriter, r *http.Request) {
		web.ServeWs(hub, w, r)
	})

	portNo := 9200

	log.Printf("ConL-Ai Server running... (portNo=%d)", portNo)
	err := http.ListenAndServe(fmt.Sprintf(":%v", portNo), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
