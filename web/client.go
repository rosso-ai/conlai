package web

import (
	"github.com/gorilla/websocket"
	"github.com/rosso-ai/conlai/conlpb"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

var repo = Repository{isEmpty: true}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  128 * 1024 * 1024,
	WriteBufferSize: 128 * 1024 * 1024,
}

func (c *Client) pullPush() {
	// pull
	doc := repo.Dequeue()
	doc.Op = "pull"
	rsp, _ := proto.Marshal(doc)
	if err := c.conn.WriteMessage(websocket.BinaryMessage, rsp); err != nil {
		log.Fatal("pull write error:", err)
		return
	}

	// push
	_, p, err := c.conn.ReadMessage()
	if err != nil || err == io.EOF {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("push read: %s", err)
		}
		return
	}

	msg := &conlpb.ConLParams{}
	err = proto.Unmarshal(p, msg)
	if err != nil {
		log.Printf("push unmarshal error: %s", err)
		return
	}
	repo.Enqueue(msg)
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()

	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("read: %s", err)
			}
			break
		}

		msg := &conlpb.ConLParams{}
		err = proto.Unmarshal(p, msg)
		if err != nil {
			log.Printf("Unmarshal: %s", err)
		}

		if msg.Op == "pull" {
			c.pullPush()

		} else if msg.Op == "update" {
			c.hub.updater <- c
			select {
			case msg := <-c.send:
				if err := c.conn.WriteMessage(websocket.BinaryMessage, msg); err != nil {
					log.Printf("update write error: %s", err)
				}
			}
		} else {
			log.Printf("cant sequence: %s", msg.Op)
		}
	}
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upGrader: %s", err)
		return
	}
	log.Print("websocket upgrade done")

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.readPump()
}
