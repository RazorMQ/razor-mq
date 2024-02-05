package razormq

import (
	"github.com/gorilla/websocket"
	"github.com/victorbetoni/razor-mq/hub"
)

type RazorMQConfig struct {
	Topics []string
	Port   int
}

type NewConsumerParams struct {
	Conn   *websocket.Conn
	Topics []string
}

type NewBrokerParams struct {
	Port   int
	Topics []string
}

type RazorMQ struct {
	Hub *hub.Hub
}

func New(config RazorMQConfig) *RazorMQ {
	h := hub.NewHub()
	for _, topic := range config.Topics {
		h.RegisterTopic(topic)
	}
	return &RazorMQ{
		Hub: h,
	}
}

func (s *RazorMQ) SetupBroker() error {
	return nil
}

func (s *RazorMQ) NewConsumer(params NewConsumerParams) error {
	return nil
}

func (s *RazorMQ) Start(port int) error {

	return nil
}
