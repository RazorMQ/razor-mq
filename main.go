package main

import (
	"github.com/victorbetoni/iris-mq/message"
	"github.com/victorbetoni/iris-mq/queue"
)

var (
	MessageQueue *queue.LinkedQueue[message.EnqueuedMessage]
)

func main() {
	MessageQueue = queue.NewLinkedQueue[message.EnqueuedMessage]()
}
