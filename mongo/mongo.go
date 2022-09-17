package mongo

import (
	"context"

	"github.com/guil95/outbox"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) outbox.Storage {
	return &mongoStorage{db}
}

func (m mongoStorage) ListAllItems(ctx context.Context) ([]outbox.Model, error) {
	cursor, err := m.db.Collection("outbox").
		Find(ctx, bson.D{{"produced", false}})

	if err != nil {
		return nil, err
	}

	var items []outbox.Model
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (m mongoStorage) UpdateItemToCheck(ctx context.Context, ids []string) error {
	_, err := m.
		db.
		Collection("outbox").
		UpdateMany(
			ctx,
			bson.M{"idempotency_id": bson.M{"$in": ids}},
			bson.D{{"$set", bson.D{{"produced", true}}}},
		)

	return err
}

func (m mongoStorage) DeleteCheckedItems(ctx context.Context) error {
	_, err := m.
		db.
		Collection("outbox").
		DeleteMany(
			ctx,
			bson.D{{"produced", true}},
		)

	return err
}

func (m mongoStorage) SaveItem(ctx context.Context, item outbox.Model) error {
	_, err := m.db.Collection("outbox").InsertOne(ctx, item)

	return err
}
