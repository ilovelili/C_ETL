package main

import (
	"database/sql"
	"dcome/query"
	"dcome/transformer"
	"etl/config"
	"etl/pipeline"
	"os"

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
	users := processors.NewSQLReader(db, query.UsersQuery())
	transformer := transformer.NewUserTransformer()
	writeCSV := processors.NewCSVWriter(os.Stdout)

	pipeline, err := pipeline.SQL_Transform_CSV(users, transformer, writeCSV)

	if err != nil {
		panic(err.Error())
	}

	err = <-pipeline.Run()
}
