package pc

import (
	"github.com/victorbetoni/razor-mq/message"
	"github.com/victorbetoni/razor-mq/queue"
)

type Consumer struct {
	Host          string
	Port          int
	MessageBuffer queue.LinkedQueue[message.EnqueuedMessage]
}
