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

// SQL_Transform_SQL defines a typical sql => transform => sql pipeline layout
func SQL_Transform_SQL(rsql *processors.SQLReader, transformer transformer.CustomTransformer, wsql *processors.SQLWriter) (pipeline *ratchet.Pipeline, err error) {
	layout, pipelineerr := ratchet.NewPipelineLayout(
		ratchet.NewPipelineStage(
			ratchet.Do(rsql).Outputs(transformer),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(transformer).Outputs(rsql),
		),
		ratchet.NewPipelineStage(
			ratchet.Do(rsql),
		),
	)

	if pipelineerr != nil {
		err = pipelineerr
	}

	// Finally, create and run the Pipeline
	pipeline = ratchet.NewBranchingPipeline(layout)
	return
}
