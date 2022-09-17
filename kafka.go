package outbox

import (
	"log"
)

type kafkaProducer struct{}

func NewKafkaProducer() Producer {
	return &kafkaProducer{}
}

func (k kafkaProducer) Produce(items []Model) error {
	log.Printf("message produced: %v", items)

	return nil
}
