package model

// OrderReceived order received by processor
type OrderReceived struct {
	SkuReceived
	UserReceived
	Item_ID      string
	Qty          int
	MCP_Order_ID string
	Created_At   string
}

// OrderTransformed order transformed by processor. Which is the basic entity for analytics
type OrderTransformed struct {
	OrderID    string
	ItemID     string
	Quantity   int
	Name       string
	Prefecture string
	City       string
	KISSkuName string
	MCPSkuName string
	Created    string
}
