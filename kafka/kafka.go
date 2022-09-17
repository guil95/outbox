package kafka

import (
	"log"

	"github.com/guil95/outbox"
)

type kafkaProducer struct{}

func NewKafkaProducer() outbox.Producer {
	return &kafkaProducer{}
}

func (k kafkaProducer) Produce(items []outbox.Model) error {
	log.Printf("message produced: %v", items)

	return nil
}
