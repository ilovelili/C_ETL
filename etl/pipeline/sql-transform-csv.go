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

// SQL_Transform_CSV defines a typical sql => transform => csv pipeline layout
func SQL_Transform_CSV(sql *processors.SQLReader, transformer transformer.CustomTransformer, writeCSV *processors.CSVWriter) (pipeline *ratchet.Pipeline, err error) {
	layout, pipelineerr := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage(
			ratchet.Do(sql).Outputs(transformer),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(transformer).Outputs(writeCSV),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(writeCSV),
		),
	)

	if pipelineerr != nil {
		err = pipelineerr
	}

	// Finally, create and run the Pipeline
	pipeline = ratchet.NewBranchingPipeline(layout)
	return
}
