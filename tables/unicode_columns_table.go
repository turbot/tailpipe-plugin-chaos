package tables

import (
	"time"

	"github.com/rs/xid"

	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

func init() {
	table.RegisterTable[*rows.UnicodeColumns, *UnicodeTestTable]()
}

const UnicodeTestTableIdentifier = "chaos_unicode_columns"

type UnicodeTestTable struct {
}

func (c *UnicodeTestTable) Identifier() string {
	return UnicodeTestTableIdentifier
}

func (c *UnicodeTestTable) GetSourceMetadata() []*table.SourceMetadata[*rows.UnicodeColumns] {
	return []*table.SourceMetadata[*rows.UnicodeColumns]{
		{
			SourceName: sources.UnicodeColumnsSourceIdentifier,
		},
	}
}

func (c *UnicodeTestTable) EnrichRow(row *rows.UnicodeColumns, sourceEnrichmentFields schema.SourceEnrichment) (*rows.UnicodeColumns, error) {
	row.CommonFields = sourceEnrichmentFields.CommonFields

	// id & Hive fields
	row.TpID = xid.New().String()
	row.TpIndex = "default"
	row.TpDate = row.Timestamp.Truncate(24 * time.Hour)

	// Timestamps
	row.TpTimestamp = row.Timestamp
	row.TpIngestTimestamp = time.Now()

	return row, nil
}
