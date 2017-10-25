package query

import "fmt"

// SQLShippingQuery query returns shipping info
func SQLShippingQuery(from string, to string) string {
	return fmt.Sprintf(`select
		t.item_id,
		t.long_item_id,
		t.tracking_id,
		t.sku,
		s.name,
		t.quantity,
		t.country,
		t.postal_code,
		t.prefecture,
		t.city,
		t.company,
		t.street1,
		t.street2,
		t.first_name,
		t.middle_name,
		t.last_name,
		t.phone,
		t.created_at
		from lms.trackings t join lms.skus s on t.sku = s.value collate utf8_unicode_ci where t.created_at between '%s' and '%s' and t.long_item_id != '';
	`, from, to)
}
