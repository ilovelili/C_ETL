package transformer

import (
	"columbus/model"
	"etl/transformer"

	"github.com/dailyburn/ratchet/data"
)

type orderTransformer struct {
	batchedOrders []model.Order
}

// NewOrderTransformer create a new order transformer
func NewOrderTransformer() transformer.CustomTransformer {
	return &orderTransformer{}
}

func (t *orderTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into slice of Order structs
	var orders = []model.Order{}

	err := data.ParseJSON(d, &orders)
	if err != nil {
		killChan <- err
	}

	// Step 2: append via pointer receiver
	t.batchedOrders = append(t.batchedOrders, orders...)
}

func (t *orderTransformer) Finish(outputChan chan data.JSON, killChan chan error) {
	var transforms []model.OrderTransfromed

	// Step 3: Loop through slice and transform data
	for _, order := range t.batchedOrders {
		transform := model.OrderTransfromed{}
		transform.Sku = order.Sku
		transform.Quantity = order.Quantity
		transform.Created = order.Created

		transforms = append(transforms, transform)
	}

	if len(transforms) > 0 {
		dd, err := data.NewJSON(transforms)
		if err != nil {
			killChan <- err
		} else {
			outputChan <- dd
		}
	}
}
