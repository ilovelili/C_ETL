package main

import (
	"database/sql"
	query "dcome/query"
	transformer "dcome/transformer"
	"etl/config"
	"os"

	"github.com/dailyburn/ratchet"
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

	layout, err := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage(
			ratchet.Do(users).Outputs(transformer),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(transformer).Outputs(writeCSV),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(writeCSV),
		),
	)

	if err != nil {
		panic(err.Error())
	}

	// Finally, create and run the Pipeline
	pipeline := ratchet.NewBranchingPipeline(layout)
	err = <-pipeline.Run()
}
