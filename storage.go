package outbox

import "context"

type Storage interface {
	ListAllItems(ctx context.Context) ([]Model, error)
	UpdateItemToCheck(ctx context.Context, id string) error
	DeleteCheckedItems(ctx context.Context) error
	SaveItem(ctx context.Context, item Model) error
}
