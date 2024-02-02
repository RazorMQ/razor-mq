package server

import (
	"github.com/victorbetoni/razor-mq/broker"
)

type NewConsumerParams struct {
	Host           string
	ReadStartIndex int
	Topics         []string
}

type NewBrokerParams struct {
	Topics []string
}

type RazorMQ struct {
	Host    string
	Brokers []*broker.Broker
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
