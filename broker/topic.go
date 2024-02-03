package broker

import (
	"time"

	"github.com/victorbetoni/razor-mq/message"
	"github.com/victorbetoni/razor-mq/queue"
)

type TopicConfig struct {
	Id           string
	MaxQueueSize int64
}

type Topic struct {
	Id           string
	CurrentIndex int64
	MaxQueueSize int64
	messageQueue *queue.LinkedQueue[message.EnqueuedMessage]
}

func New(config TopicConfig) *Topic {
	return &Topic{
		Id:           config.Id,
		CurrentIndex: 0,
		messageQueue: queue.NewLinkedQueue[message.EnqueuedMessage](),
		MaxQueueSize: config.MaxQueueSize,
	}
}

func (t *Topic) AppendMessage(msg message.Message) {
	t.messageQueue.Add(&message.EnqueuedMessage{
		ProducerHost: msg.OriginHost,
		Timestamp:    time.Now().String(),
		Index:        t.CurrentIndex,
		Data:         msg.Data,
	})
	t.CurrentIndex++
}
