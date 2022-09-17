package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/guil95/outbox"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	producer, err := kafkaProducer()
	if err != nil {
		panic(err)
	}

	mongoStorage := outbox.NewMongoStorage(mongoConnection())
	newKafkaProducer := outbox.NewKafkaProducer(producer)

	go func(str outbox.Storage, ctx context.Context) {
		for {
			_ = str.SaveItem(ctx, outbox.Model{
				IdempotencyID: uuid.NewString(),
				Message:       `{"name": "Guilherme"}`,
				Topic:         "user_created",
				Delivered:     false,
			})

			time.Sleep(time.Second * 3)
		}
	}(mongoStorage, ctx)

	ob := outbox.NewOutbox(mongoStorage, newKafkaProducer)

	ob.Listen(context.Background())
}

func mongoConnection() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", "localhost:27017"))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return client.Database("app_mongo_db")
}

func kafkaProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":        "localhost:29092",
			"delivery.timeout.ms":      600000,
			"linger.ms":                10000,
			"message.send.max.retries": 10000000,
			"batch.num.messages":       1,
			"enable.idempotence":       true,
		},
	)
}
