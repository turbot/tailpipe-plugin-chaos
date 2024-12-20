package tables

import (
	"log/slog"
	"time"

	"github.com/rs/xid"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

// init registers the table
func init() {
	// Register the table, with type parameters:
	// 1. row struct
	// 2. table config struct
	// 3. table implementation
	table.RegisterTable[*rows.AllColumns, *AllColumnsTableConfig, *AllColumnsTable]()
}

const AllColumnsTableIdentifier = "chaos_all_columns"

type AllColumnsTable struct {
}

func (c *AllColumnsTable) Identifier() string {
	return AllColumnsTableIdentifier
}

func (c *AllColumnsTable) GetSourceMetadata(_ *AllColumnsTableConfig) []*table.SourceMetadata[*rows.AllColumns] {
	return []*table.SourceMetadata[*rows.AllColumns]{
		{
			SourceName: sources.AllColumnsSourceIdentifier,
		},
	}
}

func (c *AllColumnsTable) EnrichRow(row *rows.AllColumns, _ *AllColumnsTableConfig, sourceEnrichmentFields schema.SourceEnrichment) (*rows.AllColumns, error) {
	slog.Debug(">> AllColumnsEnrichRow")

	row.CommonFields = sourceEnrichmentFields.CommonFields

	// id & Hive fields
	row.TpID = xid.New().String()
	row.TpIndex = row.TpID
	row.TpDate = row.CreatedAt.Truncate(24 * time.Hour)

	// Timestamps
	row.TpTimestamp = row.CreatedAt
	row.TpIngestTimestamp = time.Now()

	slog.Debug(">> TpIndex", "index", row.TpIndex)

	// Other Enrichment Fields
	// if row.ActorIp != "" {
	// 	row.TpSourceIP = &row.ActorIp
	// 	row.TpIps = append(row.TpIps, row.ActorIp)
	// }

	// if row.TargetId != nil {
	// 	row.TpAkas = append(row.TpAkas, *row.TargetId)
	// 	// TODO: Should row.ProcessId be added to TpAkas?
	// }

	// row.TpUsernames = append(row.TpUsernames, row.ActorHandle, row.ActorId)

	return row, nil
}
