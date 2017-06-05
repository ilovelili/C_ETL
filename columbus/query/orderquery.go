package query

import "fmt"

// SQLOrderQuery query returns order info
func SQLOrderQuery(from string, to string) string {
	return fmt.Sprintf(
		`select 
			a.display_name as name,
            a.merchant_product_id as sku,
            a.quantity, 	
			a.retail_price_list_price as itembaseprice,
			a.retail_price_discounted_price as itemdiscountprice,
			a.retail_price_base_tax as itembasetax,
			a.retail_price_shipping_list_price as itemshippingprice,
			a.retail_price_shipping_discounted_price as itemshippingdiscountprice,
			a.retail_price_shipping_tax as itemshippingtax,
			a.retail_price_total_tax as itemtotaltax,
			a.retail_price_grand_total as itemgrandtotalprice,            
            b.date_placed as created,
			b._id as orderid,
			b.line_item_count as lineitemcount,
			b.line_items_retail_price_grand_total_sum as orderprice,
			b.line_items_retail_price_discounted_price_sum as orderdiscountprice,
			b.line_items_retail_price_base_tax_sum as orderbasetax,
			b.line_items_retail_price_shipping_discounted_price_sum as ordershippingprice,
			b.line_items_retail_price_shipping_tax_sum as ordershippingtax,
			b.coupon_code as coupon,
			d.[schema] as [schema],
			d.McpSku as mcpsku
        from columbus.columbus.line_item a with(nolock)
            join columbus.columbus.sales_order b on a.sales_order_row_id = b.row_id 
			join columbus.columbus.fulfillment_sku_mapping c on a.merchant_product_id = c._id
			left join columbus..product_vpj_uniforms_products2 d on c.internal_product_id = d._id
        where 
            b.transaction_country = 'JP' and a.status != 'Canceled' and b.date_placed between '%s' and '%s'`, from, to)
}
