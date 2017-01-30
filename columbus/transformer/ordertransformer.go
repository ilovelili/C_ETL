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
	var transforms []model.OrderTransformed

	// Step 3: Loop through slice and transform data
	for _, order := range t.batchedOrders {
		transform := model.OrderTransformed{}
		transform.OrderID = order.OrderID
		transform.Sku = order.Sku
		transform.Quantity = order.Quantity
		transform.Created = order.Created

		transform.ItemBasePrice = order.ItemBasePrice
		transform.ItemDiscountPrice = order.ItemDiscountPrice
		transform.ItemBaseTax = order.ItemBaseTax
		transform.ItemShippingPrice = order.ItemShippingPrice
		transform.ItemShippingDiscountPrice = order.ItemShippingDiscountPrice
		transform.ItemShippingTax = order.ItemShippingTax
		transform.ItemTotalTax = order.ItemTotalTax
		transform.ItemGrandTotalPrice = order.ItemGrandTotalPrice

		transform.OrderPrice = order.OrderPrice
		transform.OrderDiscountPrice = order.OrderDiscountPrice
		transform.OrderBaseTax = order.OrderBaseTax
		transform.OrderShippingPrice = order.OrderShippingPrice
		transform.OrderShippingTax = order.OrderShippingTax
		transform.Schema = order.Schema

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
