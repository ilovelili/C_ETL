package query

import "fmt"

// SQLOrderQuery query returns order info
func SQLOrderQuery(from string, to string) string {
	return fmt.Sprintf(
		`select 
            a.merchant_product_id as sku,	        
            a.display_name as description,
            a.quantity, 
            CONCAT(b.billing_first_name, ' ',  b.billing_last_name) as name,
            b.billing_city as city,
            b.billing_state_or_province as prefecture,
            b.date_placed as created 
        from columbus.columbus.line_item a 
            join columbus.columbus.sales_order b on a.sales_order_row_id = b.row_id 
        where 
            b.transaction_country = 'JP' and a.status != 'Canceled' and b.date_placed between '%s' and '%s'`, from, to)
}
