package main

import (
	"github.com/victorbetoni/razor-mq/message"
	"github.com/victorbetoni/razor-mq/queue"
)

var (
	MessageQueue *queue.LinkedQueue[message.EnqueuedMessage]
)

func main() {
	MessageQueue = queue.NewLinkedQueue[message.EnqueuedMessage]()
}
