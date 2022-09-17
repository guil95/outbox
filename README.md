# Outbox
![outbox-image](./.github/images/outbox.png)

# Overview
Simple and clean implementation of outbox.

Available storages:
- MongoDB
- Mysql (in progress)

Available broker/producer:
- Kafka (in progress)

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
  "topic": "users",
  "event": "user_saved",
  "produced": false
}
```

Initialize outbox
```
mongoStorage := mongostorage.NewMongoStorage(mongoConnection())
kafkaProducer := kafka.NewKafkaProducer()

o := outbox.NewOutbox(mongoStorage, kafkaProducer)

go o.Listen(context.Background())
```

In your repository layer do you should remember to use db transaction to save to your app collection/table and outbox.
