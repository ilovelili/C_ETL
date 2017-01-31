package model

// Order order get by sql
type Order struct {
	Sku                       string
	Quantity                  int
	Created                   string
	ItemBasePrice             string
	ItemDiscountPrice         string
	ItemBaseTax               string
	ItemShippingPrice         string
	ItemShippingDiscountPrice string
	ItemShippingTax           string
	ItemTotalTax              string
	ItemGrandTotalPrice       string
	OrderID                   string
	OrderPrice                string
	OrderDiscountPrice        string
	OrderBaseTax              string
	OrderShippingPrice        string
	OrderShippingTax          string
	Coupon                    string
	Schema                    string
}

// OrderTransformed order transferred
type OrderTransformed struct {
	Sku                       string
	Quantity                  int
	Created                   string
	ItemBasePrice             string
	ItemDiscountPrice         string
	ItemBaseTax               string
	ItemShippingPrice         string
	ItemShippingDiscountPrice string
	ItemShippingTax           string
	ItemTotalTax              string
	ItemGrandTotalPrice       string
	OrderID                   string
	OrderPrice                string
	OrderDiscountPrice        string
	OrderBaseTax              string
	OrderShippingPrice        string
	OrderShippingTax          string
	Coupon                    string
	Schema                    string
}
