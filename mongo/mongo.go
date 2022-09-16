package mongo

import (
	"github.com/guil95/outbox"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) outbox.Storage {
	return &mongoStorage{db}
}

func (m mongoStorage) ListAllItems() ([]outbox.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (m mongoStorage) UpdateItemToCheck(items []outbox.Model) error {
	//TODO implement me
	panic("implement me")
}

func (m mongoStorage) DeleteCheckedItems() error {
	//TODO implement me
	panic("implement me")
}
