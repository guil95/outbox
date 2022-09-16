package outbox

import "go.uber.org/zap"

type Outbox struct {
	storage Storage
}

func (o *Outbox) Process() {
	go o.deleteCheckedItems()
	for {
		items, err := o.storage.ListAllItems()
		if err != nil {
			zap.S().Errorf("outbox: error deleting checked items %v", err)
		}
		//send to kafka
		o.updateItemToChecked(items)
	}
}

func (o *Outbox) deleteCheckedItems() {
	for {
		err := o.storage.DeleteCheckedItems()
		if err != nil {
			zap.S().Errorf("outbox: error deleting checked items %v", err)
		}
	}
}

func (o *Outbox) updateItemToChecked(items []Model) {
	err := o.storage.UpdateItemToCheck(items)
	if err != nil {
		zap.S().Errorf("outbox: error update items %v", err)
	}
}
