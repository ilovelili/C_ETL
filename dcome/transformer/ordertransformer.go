package transformer

import (
	"dcome/model"
	"etl/transformer"

	"github.com/dailyburn/ratchet/data"
)

type orderTransformer struct {
	batchedOrders []model.OrderReceived
}

// NewOrderTransformer create a new order transformer
func NewOrderTransformer() transformer.CustomTransformer {
	return &orderTransformer{}
}

func (t *orderTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into slice of OrderReceived structs
	var orders = []model.OrderReceived{}

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
		transform.OrderID = order.MCP_Order_ID
		transform.ItemID = order.Item_ID
		transform.Quantity = order.Qty
		transform.Name = order.UserReceived.First_Name + " " + order.UserReceived.Last_Name
		transform.Prefecture = order.Prefecture
		transform.City = order.UserReceived.City
		transform.KISSkuName = order.SkuReceived.Name
		transform.MCPSkuName = order.SkuReceived.Value
		transform.Created = order.Created_At

		transforms = append(transforms, transform)
	}

	// Step 4: Marshal transformed data and send to next stage
	// Write the results if more than one row/record.
	// You can change the batch size by setting loadDP.BatchSize
	if len(transforms) > 0 {
		dd, err := data.NewJSON(transforms)
		if err != nil {
			killChan <- err
		} else {
			outputChan <- dd
		}
	}
}
