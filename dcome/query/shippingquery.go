package query

import "fmt"

// SQLShippingQuery query returns shipping info
func SQLShippingQuery(from string, to string) string {
	return fmt.Sprintf(`select
		item_id,
		long_item_id,
		tracking_id,
		sku,
		quantity,
		country,
		postal_code,
		prefecture,
		city,
		company,
		street1,
		street2,
		first_name,
		middle_name,
		last_name,
		phone,
		created_at
		from lms.trackings where created_at between '%s' and '%s' and long_item_id != '';
	`, from, to)
}
