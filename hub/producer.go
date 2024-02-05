package hub

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ProducerClient struct {
	hub  *Hub
	conn *websocket.Conn
}

func NewProducerClient(hub *Hub, conn *websocket.Conn) *ProducerClient {
	return &ProducerClient{
		hub:  hub,
		conn: conn,
	}
}

func (p *ProducerClient) StartReading() {
	defer func() {
		p.hub.unregisterProducer <- p
		p.conn.Close()
	}()
	p.conn.SetReadLimit(4096)
	p.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	p.conn.SetPongHandler(func(string) error {
		p.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		return nil
	})
	for {
		_, message, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(message)
		p.hub.producerMessages <- message
	}
}
