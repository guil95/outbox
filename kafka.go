package outbox

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(producer *kafka.Producer) Producer {
	return &kafkaProducer{producer}
}

func (k kafkaProducer) Produce(items []Model, deliveredID chan<- string) error {
	go func() {
		for e := range k.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivery Sucess: %v\n", string(ev.Key))
					deliveredID <- string(ev.Key)
				}
			}
		}
	}()

	for _, item := range items {
		k.producer.ProduceChannel() <- &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &item.Topic, Partition: kafka.PartitionAny},
			Value:          []byte(item.Message),
			Key:            []byte(item.IdempotencyID),
		}
	}

	return nil
}
