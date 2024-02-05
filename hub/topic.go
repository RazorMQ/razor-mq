package hub

import (
	"encoding/json"
	"time"

	"github.com/victorbetoni/razor-mq/message"
)

type Topic struct {
	Id           string
	CurrentIndex int64
	messageChan  chan message.EnqueuedMessage
	subscribed   []*ConsumerClient
}

func NewTopic(id string) *Topic {
	return &Topic{
		Id:           id,
		CurrentIndex: 0,
		messageChan:  make(chan message.EnqueuedMessage),
		subscribed:   []*ConsumerClient{},
	}
}

func (t *Topic) StartStreaming() {
	for {
		select {
		case msg := <-t.messageChan:
			for _, c := range t.subscribed {
				b, err := json.Marshal(msg)
				if err != nil {
					continue
				}
				c.send <- b
			}
		}
	}
}

func (t *Topic) AppendMessage(msg message.Message) {
	t.messageChan <- message.EnqueuedMessage{
		ProducerHost: msg.OriginHost,
		Timestamp:    time.Now().String(),
		Index:        t.CurrentIndex,
		Data:         msg.Data,
	}
	t.CurrentIndex++
}
