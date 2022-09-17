package outbox

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type mysqlStorage struct {
	db *sqlx.DB
}

func NewMysqlStorage(db *sqlx.DB) Storage {
	return &mysqlStorage{db}
}

func (m mysqlStorage) ListAllItems(ctx context.Context) ([]Model, error) {
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

func (m mysqlStorage) SaveItem(ctx context.Context, item Model) error {
	//TODO implement me
	panic("implement me")
}
