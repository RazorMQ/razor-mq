package broker

import (
	"time"

	"github.com/victorbetoni/iris-mq/message"
	"github.com/victorbetoni/iris-mq/queue"
)

type Topic struct {
	Id           string
	CurrentIndex int
	messageQueue *queue.LinkedQueue[message.EnqueuedMessage]
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
