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

	orders := processors.NewSQLReader(db, query.OrderQuery())
	transformer := transformer.NewOrderTransformer()

	bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
	bigquery := processors.NewBigQueryWriter(bigqueryconfig, "orders")

	pipeline, err := pipeline.SQL_Transform_BigQuery(orders, transformer, bigquery)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(<-pipeline.Run())
}
