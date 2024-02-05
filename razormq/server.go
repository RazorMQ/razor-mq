package razormq

import (
	"github.com/victorbetoni/razor-mq/broker"
	"github.com/victorbetoni/razor-mq/pc"
)

type RazorMQConfig struct {
	MaxQueueSize int64
	AllTopics    []string
	Brokers      []struct {
		Port   int64
		Topics []string
	}
}

type NewConsumerParams struct {
	Host           string
	ReadStartIndex int64
	Topics         []string
}

type NewBrokerParams struct {
	Port   int
	Topics []string
}

type RazorMQ struct {
	Brokers []*broker.Broker
}

func New(config RazorMQConfig) *RazorMQ {
	razormq := &RazorMQ{}
	for _, br := range config.Brokers {
		topics := map[string]*broker.Topic{}
		for _, tp := range br.Topics {
			topics[tp] = &broker.Topic{
				Id:           tp,
				CurrentIndex: 0,
				MaxQueueSize: config.MaxQueueSize,
			}
		}
		razormq.Brokers = append(razormq.Brokers, &broker.Broker{
			Port:               int(br.Port),
			Topics:             topics,
			ConnectedConsumers: make(map[string]*pc.Consumer, 0),
		})
	}

	return razormq
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
