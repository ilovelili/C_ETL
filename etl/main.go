package main

import (
	"database/sql"
	"dcome/query"
	"dcome/transformer"
	"etl/config"
	"etl/pipeline"

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

	users := processors.NewSQLReader(db, query.UsersQuery())
	transformer := transformer.NewUserTransformer()

	bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
	bigquery := processors.NewBigQueryWriter(bigqueryconfig, "user")

	pipeline, err := pipeline.SQL_Transform_BigQuery(users, transformer, bigquery)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(<-pipeline.Run())
}
