package date_time

import (
	"log/slog"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources/date_time"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/constants"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

// init registers the table
func init() {
	// Register the table, with type parameters:
	// 1. row struct
	// 2. table config struct
	// 3. table implementation
	table.RegisterTable[*rows.DateTime, *DateTimeTable]()
}

const DateTimeTableIdentifier = "chaos_date_time"

type DateTimeTable struct {
}

func (c *DateTimeTable) Identifier() string {
	return DateTimeTableIdentifier
}

func (c *DateTimeTable) GetSourceMetadata() []*table.SourceMetadata[*rows.DateTime] {
	return []*table.SourceMetadata[*rows.DateTime]{
		{
			SourceName: date_time.DateTimeSourceIdentifier,
		},
		{
			SourceName: constants.ArtifactSourceIdentifier,
			Options: []row_source.RowSourceOption{
				artifact_source.WithArtifactExtractor(NewDateTimeExtractor()),
			},
		},
	}
}

func (c *DateTimeTable) EnrichRow(row *rows.DateTime, sourceEnrichmentFields schema.SourceEnrichment) (*rows.DateTime, error) {
	slog.Debug(">> DateTimeEnrichRow")

	row.CommonFields = sourceEnrichmentFields.CommonFields

	// id & Hive fields
	row.TpID = row.Id
	row.TpIndex = row.TpID
	row.TpDate = row.Timestamp.Truncate(24 * time.Hour)

	// Timestamps
	row.TpTimestamp = row.Timestamp
	row.TpIngestTimestamp = time.Now()

	return row, nil
}
