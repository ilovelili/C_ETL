package main

import (
	"database/sql"
	"dcome/query"
	"dcome/transformer"
	"etl/config"
	"etl/pipeline"
	"flag"
	"os"
	"time"

	"github.com/dailyburn/ratchet/processors"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	date = flag.String("date", "", "object date")
)

func main() {
	flag.Parse()

	if *date == "" {
		*date = time.Now().Format("2006-01-02")
	}

	config := config.GetConfig()
	db, err := sql.Open(config.Client, config.GetConnectionString())
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	orders := processors.NewSQLReader(db, query.SQLOrderQuery(*date))
	transformer := transformer.NewOrderTransformer()

	bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
	bigquery := processors.NewBigQueryWriter(bigqueryconfig, config.DataTable)

	pipeline, err := pipeline.SQL_Transform_BigQuery(orders, transformer, bigquery)

	if err != nil {
		panic(err.Error())
	}

	<-pipeline.Run()
	os.Exit(0)
}
