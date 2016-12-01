package transformer

import (
	model "dcome/model"

	"github.com/dailyburn/ratchet/data"
)

type userTransformer struct {
	batchedUsers []model.UserReceived
}

// NewUserTransformer create a new user transformer
func NewUserTransformer() *userTransformer {
	return &userTransformer{}
}

func (t *userTransformer) ProcessData(d data.JSON, outputChan chan data.JSON, killChan chan error) {
	// Step 1: Unmarshal json into slice of UserReceived structs
	var users = []model.UserReceived{}

	err := data.ParseJSON(d, &users)
	if err != nil {
		killChan <- err
	}

	// Step 2: append via pointer receiver
	t.batchedUsers = append(t.batchedUsers, users...)
}

func (t *userTransformer) Finish(outputChan chan data.JSON, killChan chan error) {
	var transforms []model.UserTransformed

	// Step 3: Loop through slice and transform data
	for _, user := range t.batchedUsers {
		transform := model.UserTransformed{}
		transform.Prefecture = user.Prefecture
		transform.City = user.City
		transform.Name = user.First_Name + " " + user.Last_Name
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
