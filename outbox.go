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
	go o.deleteCheckedItems(ctx)

	itemsFound := make(chan []Model)
	deliveredChan := make(chan string)

	go func(itemsFound chan []Model, deliveredChan chan string) {
		for {
			select {
			case idempotencyID := <-deliveredChan:
				o.updateItemToChecked(ctx, idempotencyID)
			case i := <-itemsFound:
				if err := o.producer.Produce(i, deliveredChan); err != nil {
					log.Printf("error to produce message %v", err)
					continue
				}
			}
		}
	}(itemsFound, deliveredChan)

	for {
		items, err := o.storage.ListAllItems(ctx)
		if err != nil {
			log.Printf("outbox: error deleting checked items %v", err)
		}

		itemsFound <- items

		time.Sleep(time.Second * 2)
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

func (o *outbox) updateItemToChecked(ctx context.Context, idempotencyID string) {
	err := o.storage.UpdateItemToCheck(ctx, idempotencyID)
	if err != nil {
		log.Printf("outbox: error update items %v", err)
	}
}
