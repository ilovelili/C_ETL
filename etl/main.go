package main

import (
	"database/sql"
	"dcome/query"
	"dcome/transformer"
	"etl/config"
	"etl/pipeline"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/processors"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/japanese"
)

var (
	mode = flag.String("mode", "", "define run mode")
)

func main() {
	flag.Parse()

	// export data to cloud
	if *mode == "export" {
		dump()
		os.Exit(0)
	}

	if *mode == "import" {
		read()
		os.Exit(0)
	}

	fmt.Println("Mode unassigned. Exit")
	os.Exit(0)
}

// on-premised to cloud
func dump() {
	config := config.GetConfig()
	db, err := sql.Open(config.Client, config.GetConnectionString())
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	orders := processors.NewSQLReader(db, query.SQLOrderQuery())
	transformer := transformer.NewOrderTransformer()

	bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
	bigquery := processors.NewBigQueryWriter(bigqueryconfig, config.DataTable)

	pipeline, err := pipeline.SQL_Transform_BigQuery(orders, transformer, bigquery)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(<-pipeline.Run())
}

// cloud to on-premises
func read() {
	config := config.GetConfig()

	bigqueryconfig := &processors.BigQueryConfig{JsonPemPath: config.JsonPemPath, ProjectID: config.ProjectID, DatasetID: config.DatasetID}
	bigquerysku := processors.NewBigQueryReader(bigqueryconfig, query.BigQuerySkuQuery())
	bigquerydaily := processors.NewBigQueryReader(bigqueryconfig, query.BigQueryDailyQuery())

	// use shift_jis
	encoder := japanese.ShiftJIS.NewEncoder()
	sku, _ := os.Create(config.CSVSavePath + "sku" + time.Now().Format("20060102") + ".csv")
	daily, _ := os.Create(config.CSVSavePath + "daily" + time.Now().Format("20060102") + ".csv")
	defer sku.Close()
	defer daily.Close()

	writeSkuCSV := processors.NewCSVWriter(encoder.Writer(sku))
	writeDailyCSV := processors.NewCSVWriter(encoder.Writer(daily))

	layout, err := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage(
			ratchet.Do(bigquerysku).Outputs(writeSkuCSV),
			ratchet.Do(bigquerydaily).Outputs(writeDailyCSV),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(writeSkuCSV),
			ratchet.Do(writeDailyCSV),
		),
	)

	if err != nil {
		panic(err.Error())
	}

	// Finally, create and run the Pipeline
	pipeline := ratchet.NewBranchingPipeline(layout)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(<-pipeline.Run())
}
