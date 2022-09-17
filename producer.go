package outbox

type Producer interface {
	Produce(items []Model, deliveredID chan<- string) error
}
