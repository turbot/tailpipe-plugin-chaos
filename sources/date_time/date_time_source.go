package date_time

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources"
	"github.com/turbot/tailpipe-plugin-chaos/sources/all_columns"
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const DateTimeSourceIdentifier = "chaos_date_time"

// DateTimeSource source is responsible for collecting logs
type DateTimeSource struct {
	// row_source.RowSourceImpl[*DateTimeSourceConfig]
	row_source.RowSourceImpl[*DateTimeSourceConfig, *config.ChaosConnection]
}

func (s *DateTimeSource) Init(ctx context.Context, params *row_source.RowSourceParams, opts ...row_source.RowSourceOption) error {
	// set the collection state ctor
	s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	return s.RowSourceImpl.Init(ctx, params, opts...)
}

func (s *DateTimeSource) Identifier() string {
	return DateTimeSourceIdentifier
}

func (s *DateTimeSource) Collect(ctx context.Context) error {

	// populate enrichment fields the source is aware of
	// - in this case the connection
	sourceName := all_columns.AllColumnsSourceIdentifier
	sourceEnrichmentFields := &schema.SourceEnrichment{
		CommonFields: schema.CommonFields{
			TpSourceName: &sourceName,
			TpSourceType: DateTimeSourceIdentifier, // TODO: review this
		},
	}

	slog.Debug(">> Collecting data from source")

	// set the current time to 2006-01-02 15:04:05 +0000 UTC
	currentTime := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	for i := 1; i <= 100; i++ {
		// populate the row data
		rowData := s.populateDateTime(currentTime)

		row := &types.RowData{Data: rowData, SourceEnrichment: sourceEnrichmentFields}
		slog.Debug(">> Sending row to plugin", "row", row)
		if err := s.OnRow(ctx, row); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}

		// Increment the time by 4 days
		currentTime = currentTime.Add(4 * 24 * time.Hour)
	}

	return nil
}

func (s *DateTimeSource) populateDateTime(currentTime time.Time) *rows.DateTime {
	return &rows.DateTime{
		Id:        sources.RandomString(10),
		Timestamp: currentTime,
	}
}
