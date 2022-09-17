package outbox

type Model struct {
	IdempotencyID string `json:"idempotency_id" bson:"idempotency_id"`
	Message       string `json:"message" bson:"message"`
	Topic         string `json:"topic" bson:"topic"`
	Event         string `json:"event" bson:"event"`
	Produced      bool   `json:"produced" bson:"produced"`
}
