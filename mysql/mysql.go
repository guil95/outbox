package mysql

import (
	"context"

	"github.com/guil95/outbox"
	"github.com/jmoiron/sqlx"
)

type mysqlStorage struct {
	db *sqlx.DB
}

func NewMysqlStorage(db *sqlx.DB) outbox.Storage {
	return &mysqlStorage{db}
}

func (m mysqlStorage) ListAllItems(ctx context.Context) ([]outbox.Model, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlStorage) UpdateItemToCheck(ctx context.Context, ids []string) error {
	//TODO implement me
	panic("implement me")
}

func (m mysqlStorage) DeleteCheckedItems(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (m mysqlStorage) SaveItem(ctx context.Context, item outbox.Model) error {
	//TODO implement me
	panic("implement me")
}
