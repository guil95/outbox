package outbox

type Model struct {
	ID      string `json:"id" bson:"id"`
	Payload string `json:"payload" bson:"payload"`
	Topic   string `json:"topic" bson:"topic"`
	Event   string `json:"event" bson:"event"`
	Checked bool   `json:"checked" bson:"checked"`
}
