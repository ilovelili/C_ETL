package query

import "fmt"

// UsersQuery query returns user info
func UsersQuery() string {
	return fmt.Sprintf(`select * FROM lms.addresses`)
}
