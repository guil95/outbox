# Outbox
![outbox-image](./.github/images/outbox.png)

# Overview
Simple and clean implementation of outbox.

Available storages:
- MongoDB
- Mysql (in progress)

Available broker/producer:
- Kafka 

# Usage

Create a collection on mongodb with this fields

```json
{
  "idempotency_id": "58452f68-705b-4b2e-8685-fc929e750588",
  "message": {
    "id": "58452f68-705b-4b2e-8685-fc929e750588",
    "name": "Guilherme",
    "age": 27
  },
  "topic": "user_created"
}
```

Initialize outbox
```
mongoStorage := mongostorage.NewMongoStorage(mongoConnection())
kafkaProducer := kafka.NewProducer(kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers":        "localhost:29092",
			"delivery.timeout.ms":      600000,
			"linger.ms":                10000,
			"message.send.max.retries": 10000000,
			"batch.num.messages":       1,
			"enable.idempotence":       true,
		},
	))

o := outbox.NewOutbox(mongoStorage, kafkaProducer)

go o.Listen(context.Background())
```

In your storage layer do you should remember to use db transaction to save to your app collection/table and outbox.

# Example
Use the docker-compose inside example_mongo to up dependencies and run the example

> **It's recommend use the configuration** `batch.num.messages` **with value** `1` **to will don't have problem with duplicate messages, also validate idempotency_id on consumer**
