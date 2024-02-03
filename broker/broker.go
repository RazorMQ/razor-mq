package broker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/victorbetoni/razor-mq/message"
	"github.com/victorbetoni/razor-mq/pc"
)

const (
	MAX_MESSAGE_SIZE = 2048
)

type brokerMessage struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

type Broker struct {
	Port               int
	msgChan            chan message.Message
	Topics             map[string]*Topic
	ConnectedConsumers map[string]*pc.Consumer
}

func (b *Broker) HandleMessage(msg message.Message) error {
	topic, ok := b.Topics[msg.Topic]
	if !ok {
		return fmt.Errorf("topic not found: %s", msg.Topic)
	}
	topic.AppendMessage(msg)
	return nil
}

func (b *Broker) Subscribe(host string, port int, topics []string) {

	pc.Consumer

}

func (b *Broker) Listen() {
	b.msgChan = make(chan message.Message)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", b.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		defer listener.Close()
		for {
			connection, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go func(c net.Conn) {
				io.Copy(c, c)
				defer c.Close()

				buf := make([]byte, MAX_MESSAGE_SIZE)
				_, err := c.Read(buf)

				if err != nil {
					log.Fatal(err)
				}

				data := &brokerMessage{}
				if err := json.Unmarshal(buf, data); err != nil {
					log.Fatal("broker message with wrong format")
					c.Close()
					return
				}

				msg := message.Message{
					Data:       []byte(data.Data),
					OriginHost: c.RemoteAddr().String(),
					Topic:      data.Topic,
				}
				b.msgChan <- msg
			}(connection)
		}
	}()

	go func() {
		for {
			msg := <-b.msgChan
			b.Topics[msg.Topic].AppendMessage(msg)
		}
	}()
}
