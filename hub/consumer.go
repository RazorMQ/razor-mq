package hub

import (
	"time"

	"github.com/gorilla/websocket"
)

type ConsumerClient struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewConsumerClient(hub *Hub, topics []string, conn *websocket.Conn) *ConsumerClient {
	return &ConsumerClient{
		hub:  hub,
		conn: conn,
		send: make(chan []byte),
	}
}

func (c *ConsumerClient) Stream() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
