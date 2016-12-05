package query

import "fmt"

// OrderQuery query returns order info
func OrderQuery() string {
	return fmt.Sprintf(
		`select * FROM lms.line_items items 
		join lms.orders orders on orders.id = items.order_id 
		join lms.addresses addresses on addresses.id = items.address_id 
		join lms.skus skus on skus.id = items.sku_id where items.line_status_id != 1`)
}
