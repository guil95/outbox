package outbox

type Model struct {
	IdempotencyID string `json:"idempotency_id" bson:"idempotency_id"`
	Message       string `json:"message" bson:"message"`
	Topic         string `json:"topic" bson:"topic"`
	Delivered     bool   `json:"delivered" bson:"delivered"`
}
