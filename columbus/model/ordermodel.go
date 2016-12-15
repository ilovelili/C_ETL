package model

// Order order get by sql
type Order struct {
	Name        string
	Sku         string
	Quantity    int
	Description string
	City        string
	Prefecture  string
	Created     string
}

// OrderTransfromed order transferred
type OrderTransfromed struct {
	Sku      string
	Quantity int
	Created  string
}
