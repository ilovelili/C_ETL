package main

import (
	"database/sql"
	"dcome"
	"etl/config"

	"fmt"

	"github.com/dailyburn/ratchet/processors"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	config := config.GetConfig()
	db, err := sql.Open(config.Client, config.GetConnectionString())
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	// First initalize the DataProcessors
	read := processors.NewSQLReader(db, dcome.UsersQuery())

	fmt.Println(read)
}
