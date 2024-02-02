package server

import (
	"github.com/victorbetoni/iris-mq/broker"
)

type NewConsumerParams struct {
	Host           string
	ReadStartIndex int
	Topics         []string
}

type NewBrokerParams struct {
	Topics []string
}

type IrisMQ struct {
	Host    string
	Brokers []*broker.Broker
}

func (s *IrisMQ) SetupBroker() error {
	return nil
}

func (s *IrisMQ) NewConsumer(params NewConsumerParams) error {
	return nil
}

func (s *IrisMQ) Start(port int) error {

	return nil
}
