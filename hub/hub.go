package hub

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/victorbetoni/razor-mq/razormq"
)

type Handshake struct {
	ApplicationType string   `json:"application_type"`
	ReadStartIndex  int64    `json:"start_index"`
	Topics          []string `json:"topics"`
	Host            string
	Port            int
}

type Hub struct {
	razorMq         *razormq.RazorMQ
	consumerHubPort int
	consumerHubChan chan Handshake
}

func NewHub(razorMq *razormq.RazorMQ, port int) *Hub {
	return &Hub{
		razorMq:         razorMq,
		consumerHubPort: port,
		consumerHubChan: make(chan Handshake),
	}
}

func (h *Hub) OpenConsumerHub() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", h.consumerHubPort))
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()
		for {
			connection, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
				continue
			}
			go func(c net.Conn) {
				io.Copy(c, c)
				defer c.Close()

				buf := make([]byte, 1024)
				_, err := c.Read(buf)
				if err != nil {
					log.Fatal(err)
				}

				data := &Handshake{}
				if err := json.Unmarshal(buf, data); err != nil {
					log.Fatal("handshake message with wrong format")
					c.Close()
					return
				}
				data.Host = c.RemoteAddr().String()
				h.consumerHubChan <- *data
			}(connection)
		}
	}()

	go func() {
		for {
			handshake := <-h.consumerHubChan
			if strings.EqualFold(handshake.ApplicationType, "consumer") {
				h.razorMq.NewConsumer(razormq.NewConsumerParams{
					Host:           handshake.Host,
					ReadStartIndex: handshake.ReadStartIndex,
					Topics:         handshake.Topics,
				})
			}
		}
	}()
}
