package query

import "fmt"

// SQLOrderQuery query returns order info
func SQLOrderQuery(date string) string {
	return fmt.Sprintf(
		`select 
		items.item_id, 
		items.qty, 
		addresses.country, 
		addresses.first_name, 
		addresses.last_name, 
		addresses.prefecture, 
		addresses.city, 
		orders.mcp_order_id, 
		orders.created_at, 
		skus.name, 
		skus.value 
		FROM lms.line_items items 
		join lms.orders orders on orders.id = items.order_id 
		join lms.addresses addresses on addresses.id = items.address_id 
		join lms.skus skus on skus.id = items.sku_id where items.line_status_id != 1 and DATE(orders.created_at)=%s`, date)
}
