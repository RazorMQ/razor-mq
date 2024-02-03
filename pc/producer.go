package pc

import "github.com/victorbetoni/razor-mq/message"

type Producer struct {
	Port int 
	Host string
}

func (p *Producer) HandleMessage(msg message.Message) {

}
