package pc

import (
	"fmt"
	"net"
)

type Consumer struct {
	Host       string
	Port       int
	Connection *net.Conn
}

func (c *Consumer) OpenConnection() (*net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, err
	}
	c.Connection = &conn
	return &conn, nil
}
