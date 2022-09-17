package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/guil95/outbox"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	mongoStorage := outbox.NewMongoStorage(mongoConnection())
	kafkaProducer := outbox.NewKafkaProducer()

	go func(str outbox.Storage, ctx context.Context) {
		for {
			_ = str.SaveItem(ctx, outbox.Model{
				IdempotencyID: uuid.NewString(),
				Message:       `{"name": "Guilherme"}`,
				Topic:         "users_ms",
				Event:         "user_saved",
				Produced:      false,
			})

			time.Sleep(time.Second * 2)
		}
	}(mongoStorage, ctx)

	ob := outbox.NewOutbox(mongoStorage, kafkaProducer)

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

	return client.Database(os.Getenv("MONGO_OUTBOX_DB"))
}
