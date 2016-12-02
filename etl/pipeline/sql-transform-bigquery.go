// Package pipeline defines common static patterns of etl pipeline like:
// mysql -> custom transform -> csv
// mysql -> aws elasticmapreduce -> s3 -> mysql
// mssql -> google big query -> mssql
// any other common pattern
package pipeline

import (
	"etl/transformer"

	"github.com/dailyburn/ratchet"
	"github.com/dailyburn/ratchet/processors"
)

// SQL_Transform_BigQuery defines a typical sql => transform => bigquery pipeline layout
func SQL_Transform_BigQuery(sql *processors.SQLReader, transformer transformer.CustomTransformer, bigQuery *processors.BigQueryWriter) (pipeline *ratchet.Pipeline, err error) {
	layout, pipelineerr := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage(
			ratchet.Do(sql).Outputs(transformer),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(transformer).Outputs(bigQuery),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(bigQuery),
		),
	)

	if pipelineerr != nil {
		err = pipelineerr
	}

	// Finally, create and run the Pipeline
	pipeline = ratchet.NewBranchingPipeline(layout)
	return
}
