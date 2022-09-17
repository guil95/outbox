package outbox

import (
	"context"
	"log"
	"time"
)

type Outbox interface {
	Listen(ctx context.Context)
}

type outbox struct {
	storage  Storage
	producer Producer
}

func NewOutbox(storage Storage, producer Producer) Outbox {
	return &outbox{storage, producer}
}

func (o *outbox) Listen(ctx context.Context) {
	log.Println("initialize outbox process")

	go o.deleteCheckedItems(ctx)
	for {
		items, err := o.storage.ListAllItems(ctx)

		log.Printf("items found %v", items)

		if err != nil {
			log.Printf("outbox: error deleting checked items %v", err)
		}

		err = o.producer.Produce(items)
		if err != nil {
			log.Printf("error to produce message %v", err)
			continue
		}

		o.updateItemToChecked(ctx, items)

		time.Sleep(time.Second * 1)
	}
}

func (o *outbox) deleteCheckedItems(ctx context.Context) {
	for {
		err := o.storage.DeleteCheckedItems(ctx)
		if err != nil {
			log.Printf("outbox: error deleting checked items %v", err)
		}

		time.Sleep(time.Second * 1)
	}
}

func (o *outbox) updateItemToChecked(ctx context.Context, items []Model) {
	var ids []string

	for _, item := range items {
		ids = append(ids, item.IdempotencyID)
	}

	if len(ids) == 0 {
		return
	}

	err := o.storage.UpdateItemToCheck(ctx, ids)
	if err != nil {
		log.Printf("outbox: error update items %v", err)
	}
}
