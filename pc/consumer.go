package pc

import (
	"github.com/victorbetoni/iris-mq/message"
	"github.com/victorbetoni/iris-mq/queue"
)

type Consumer struct {
	Host          string
	Port          int
	MessageBuffer queue.LinkedQueue[message.EnqueuedMessage]
}
