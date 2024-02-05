package hub

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/RazorMQ/razor-mq/message"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Handshake struct {
	ApplicationType string   `json:"application_type"`
	ReadStartIndex  int64    `json:"start_index"`
	Topics          []string `json:"topics"`
	Host            string
	Port            int
}

type Hub struct {
	registeredTopics    map[string]*Topic
	registeredProducers map[*ProducerClient]bool
	registeredConsumers map[*ConsumerClient]bool
	producerMessages    chan []byte
	unregisterProducer  chan *ProducerClient
	unregisterConsumer  chan *ConsumerClient
	registerProducer    chan *ProducerClient
	registerConsumer    chan *ConsumerClient
}

func NewHub() *Hub {
	return &Hub{
		registeredTopics:    make(map[string]*Topic),
		registeredProducers: make(map[*ProducerClient]bool),
		registeredConsumers: make(map[*ConsumerClient]bool),
		producerMessages:    make(chan []byte),
		registerProducer:    make(chan *ProducerClient),
		registerConsumer:    make(chan *ConsumerClient),
		unregisterProducer:  make(chan *ProducerClient),
		unregisterConsumer:  make(chan *ConsumerClient),
	}
}

func (h *Hub) RegisterTopic(id string) {
	h.registeredTopics[id] = NewTopic(id)
}

func (h *Hub) Start() {
	for {
		select {
		case producer := <-h.registerProducer:
			h.registeredProducers[producer] = true
			go producer.StartReading()
		case consumer := <-h.registerConsumer:
			h.registeredConsumers[consumer] = true
			go consumer.Stream()
		case producer := <-h.unregisterProducer:
			h.registeredProducers[producer] = false
		case consumer := <-h.unregisterConsumer:
			h.registeredConsumers[consumer] = false
		case message := <-h.producerMessages:
			h.handleMessage(message)
		}
	}
}

func (h *Hub) handleMessage(msg []byte) {
	message := &message.Message{}

	if err := json.Unmarshal(msg, message); err != nil {
		return
	}

	topic, ok := h.registeredTopics[message.Topic]
	if !ok {
		return
	}

	topic.AppendMessage(*message)
}

func (h *Hub) HandlePeer(w http.ResponseWriter, r *http.Request) {
	t, ok := r.Header[http.CanonicalHeaderKey("RazorMQ-Application-Type")]
	if !ok {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if strings.EqualFold(t[0], "consumer") {
		client := &ConsumerClient{
			hub:  h,
			conn: conn,
			send: make(chan []byte),
		}
		h.registerConsumer <- client
	}

	if strings.EqualFold(t[0], "producer") {
		client := &ProducerClient{
			hub:  h,
			conn: conn,
		}
		h.registerProducer <- client
	}

}
