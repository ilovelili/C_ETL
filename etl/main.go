package main

import (
	"database/sql"
	dcomequery "dcome/query"
	dcometransformer "dcome/transformer"
	"strings"

	columbusquery "columbus/query"
	columbustransformer "columbus/transformer"

	"etl/config"
	"etl/pipeline"
	"flag"
	"time"

	"etl/transformer"

	"github.com/dailyburn/ratchet/processors"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	from   = flag.String("from", time.Now().Format("2006-01-02") /*today*/, "date from")
	to     = flag.String("to", time.Now().AddDate(0, 0, 1).Format("2006-01-02") /*yesterday*/, "date to")
	domain = flag.String("domain", "columbus", "domain to select. DCOME or Columbus")
)

func main() {
	flag.Parse()

	config := config.GetConfig()
	db, err := sql.Open(config.Client, config.GetConnectionString())
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	var orders *processors.SQLReader
	var transformer transformer.CustomTransformer
	var bigquery *processors.BigQueryWriter

	if strings.ToLower(*domain) == "dcome" {
		orders = processors.NewSQLReader(db, dcomequery.SQLOrderQuery(*from, *to))
		transformer = dcometransformer.NewOrderTransformer()

		bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
		bigquery = processors.NewBigQueryWriter(bigqueryconfig, config.DataTable)

	} else if strings.ToLower(*domain) == "columbus" {
		orders = processors.NewSQLReader(db, columbusquery.SQLOrderQuery(*from, *to))
		transformer = columbustransformer.NewOrderTransformer()

		bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
		bigquery = processors.NewBigQueryWriter(bigqueryconfig, config.DataTable)
	}

	pipeline, err := pipeline.SQL_Transform_BigQuery(orders, transformer, bigquery)
	if err != nil {
		panic(err.Error())
	}

	<-pipeline.Run()
	// os.Exit(0)
}
