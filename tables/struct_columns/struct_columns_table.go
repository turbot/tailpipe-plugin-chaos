package struct_columns

import (
	"time"

	"github.com/rs/xid"

	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources/struct_columns"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

// init registers the table
func init() {
	// Register the table, with type parameters:
	// 1. row struct
	// 2. table config struct
	// 3. table implementation
	table.RegisterTable[*rows.StructColumns, *StructColumnsTable]()
}

const StructColumnsTableIdentifier = "chaos_struct_columns"

type StructColumnsTable struct {
}

func (c *StructColumnsTable) Identifier() string {
	return StructColumnsTableIdentifier
}

func (c *StructColumnsTable) GetSourceMetadata() ([]*table.SourceMetadata[*rows.StructColumns], error) {
	return []*table.SourceMetadata[*rows.StructColumns]{
		{
			SourceName: struct_columns.StructColumnsSourceIdentifier,
		},
	}, nil
}

func (c *StructColumnsTable) EnrichRow(row *rows.StructColumns, sourceEnrichmentFields schema.SourceEnrichment) (*rows.StructColumns, error) {
	row.CommonFields = sourceEnrichmentFields.CommonFields

	row.TpID = xid.New().String()
	row.TpIndex = row.Timestamp.Format("2006_01_02_06_05_04")
	row.TpTimestamp = row.Timestamp
	row.TpDate = row.Timestamp.Truncate(24 * time.Hour)
	row.TpIngestTimestamp = time.Now()

	return row, nil
}
