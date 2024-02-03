package pc

import (
	"fmt"
	"log"
	"net"

	"github.com/victorbetoni/razor-mq/message"
)

type Consumer struct {
	Host       string
	Port       int
	Connection *net.Conn
}

func (c *Consumer) HandleMessage(message.EnqueuedMessage) {
	if c.Connection == nil {
		if _, err := c.OpenConnection(); err != nil {
			log.Fatalf("couldnt stablish connection with consumer %s:%d: %s", c.Host, c.Port, err.Error())
			return
		}
	}
}

func (c *Consumer) OpenConnection() (*net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, err
	}
	c.Connection = &conn
	return &conn, nil
}
