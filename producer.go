package outbox

type Producer interface {
	Produce(items []Model) error
}
