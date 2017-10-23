package query

import "fmt"

// SQLShippingQuery query returns shipping info
func SQLShippingQuery(from string, to string) string {
	return fmt.Sprintf(`SELECT * FROM lms.trackings where created_at between '%s' and '%s';`, from, to)
}
