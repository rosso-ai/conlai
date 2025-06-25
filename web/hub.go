package web

import (
	"github.com/rosso-ai/conlai/conlpb"
	"google.golang.org/protobuf/proto"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client

	updateWaiting map[*Client]bool
	updater       chan *Client
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),

		updateWaiting: make(map[*Client]bool),
		updater:       make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.updateWaiting[client] = false

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}

		case client := <-h.updater:
			h.updateWaiting[client] = true
			allReceived := true
			for client := range h.clients {
				if h.updateWaiting[client] == false {
					allReceived = false
				}
			}

			if allReceived {
				msg := &conlpb.ConLParams{Op: "update"}
				rsp, _ := proto.Marshal(msg)
				for client := range h.clients {
					client.send <- rsp
					h.updateWaiting[client] = false
				}
			}
		}
	}
}
