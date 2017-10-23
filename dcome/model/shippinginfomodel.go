package model

// ShippingInfoReceived shipping info received by processor. Which is the basic entity for analytics
type ShippingInfoReceived struct {
	Item_ID      string
	Long_Item_ID string
	Tracking_ID  string
	Sku          string
	Quantity     int
	Country      string
	Postal_Code  string
	Prefecture   string
	City         string
	Company      string
	Street1      string
	Street2      string
	First_Name   string
	Middle_Name  string
	Last_Name    string
	Phone        string
	Created_At   string
}

// ShippingInfoTransformed shipping info transformed by processor. Which is the basic entity for analytics
type ShippingInfoTransformed struct {
	ItemID     string
	LongItemID string
	TrackingID string
	Sku        string
	Quantity   int
	Country    string
	PostalCode string
	Address    string
	Company    string
	Name       string
	Phone      string
	Created    string
}
