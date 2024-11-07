package tables

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/rs/xid"
	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/helpers"
	"github.com/turbot/tailpipe-plugin-sdk/parse"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

const AllColumnsTableIdentifier = "chaos_all_columns"

type AllColumnsTable struct {
	// all tables must embed table.TableImpl
	table.TableImpl[*rows.AllColumns, *AllColumnsTableConfig, *config.ChaosConnection]
}

func NewAllColumnsTable() table.Table {
	return &AllColumnsTable{}
}

func (c *AllColumnsTable) Identifier() string {
	return AllColumnsTableIdentifier
}

// GetRowSchema implements Table
// return an instance of the row struct
func (c *AllColumnsTable) GetRowSchema() any {
	return rows.AllColumns{}
}

func (c *AllColumnsTable) GetConfigSchema() parse.Config {
	return &AllColumnsTableConfig{}
}

func (c *AllColumnsTable) EnrichRow(row *rows.AllColumns, sourceEnrichmentFields *enrichment.CommonFields) (*rows.AllColumns, error) {
	slog.Debug(">> AllColumnsEnrichRow")
	// we expect sourceEnrichmentFields to be set
	if sourceEnrichmentFields == nil {
		return nil, fmt.Errorf("AllColumns EnrichRow called with nil sourceEnrichmentFields")
	}
	// we expect name to be set by the Source
	if sourceEnrichmentFields.TpSourceName == "" {
		return nil, fmt.Errorf("AllColumnsTable EnrichRow called with TpSourceName unset in sourceEnrichmentFields")
	}

	row.CommonFields = *sourceEnrichmentFields

	// id & Hive fields
	row.TpID = xid.New().String()
	row.TpIndex = c.Config.Index
	row.TpDate = row.CreatedAt.Format("2006-01-02")

	// Timestamps
	row.TpTimestamp = helpers.UnixMillis(row.CreatedAt.UnixNano() / int64(time.Millisecond))
	row.TpIngestTimestamp = helpers.UnixMillis(time.Now().UnixNano() / int64(time.Millisecond))

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
