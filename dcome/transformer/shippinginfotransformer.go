package transformer

import (
	"dcome/model"
	"etl/transformer"

	"github.com/dailyburn/ratchet/data"
)

type shippingInfoTransformer struct {
	batchedShippingInfo []model.ShippingInfoReceived
}

// NewShippingInfoTransformer create a new shipping info transformer
func NewShippingInfoTransformer() transformer.CustomTransformer {
	return &shippingInfoTransformer{}
}

func (t *shippingInfoTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into slice of ShippingInfoReceived structs
	var shippiginfos = []model.ShippingInfoReceived{}

	err := data.ParseJSON(d, &shippiginfos)
	if err != nil {
		killChan <- err
	}

	// Step 2: append via pointer receiver
	t.batchedShippingInfo = append(t.batchedShippingInfo, shippiginfos...)
}

func (t *shippingInfoTransformer) Finish(outputChan chan data.JSON, killChan chan error) {
	var transforms []model.ShippingInfoTransformed

	// Step 3: Loop through slice and transform data
	for _, shippinginfo := range t.batchedShippingInfo {
		transform := model.ShippingInfoTransformed{}
		transform.ItemID = shippinginfo.Item_ID
		transform.LongItemID = shippinginfo.Long_Item_ID
		transform.TrackingID = shippinginfo.Tracking_ID
		transform.Quantity = shippinginfo.Quantity
		transform.Country = shippinginfo.Country
		transform.PostalCode = shippinginfo.Postal_Code
		transform.Address = shippinginfo.Prefecture + " " + shippinginfo.City + " " + shippinginfo.Street1 + " " + shippinginfo.Street2
		transform.Company = shippinginfo.Company
		transform.Name = shippinginfo.First_Name + shippinginfo.Middle_Name + shippinginfo.Last_Name
		transform.Phone = shippinginfo.Phone
		transform.Created = shippinginfo.Created_At

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
