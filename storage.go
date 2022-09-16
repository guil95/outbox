package outbox

type Storage interface {
	ListAllItems() ([]Model, error)
	UpdateItemToCheck(items []Model) error
	DeleteCheckedItems() error
}
