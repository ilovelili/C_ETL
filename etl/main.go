package main

import (
	"database/sql"
	"dcome"
	"etl/config"

	"fmt"

	"github.com/dailyburn/ratchet/processors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	client           string
	connectionstring string
)

func init() {
	config := config.GetConfig()

	client = config.OnPremisesConfig.DatabaseConfig.Client
	connectionstring = config.OnPremisesConfig.DatabaseConfig.User + ":" +
		config.OnPremisesConfig.DatabaseConfig.Password + "@/" +
		config.OnPremisesConfig.DatabaseConfig.Database
}

func main() {
	db, err := sql.Open(client, connectionstring)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	// First initalize the DataProcessors
	read := processors.NewSQLReader(db, dcome.UsersQuery())

	fmt.Println(read)
}
