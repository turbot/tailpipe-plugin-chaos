package sources

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
	"github.com/turbot/tailpipe-plugin-sdk/config_data"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/types"
	"golang.org/x/exp/rand"
)

const AllColumnsSourceIdentifier = "chaos_all_columns"

// init function should register the source
func init() {
	row_source.RegisterRowSource[*AllColumnsSource]()
}

// AllColumnsSource source is responsible for collecting audit logs from Turbot Pipes API
type AllColumnsSource struct {
	// row_source.RowSourceImpl[*AllColumnsSourceConfig]
	row_source.RowSourceImpl[*AllColumnsSourceConfig, *config.ChaosConnection]
}

func (s *AllColumnsSource) Init(ctx context.Context, configData config_data.ConfigData, connectionData config_data.ConfigData, opts ...row_source.RowSourceOption) error {
	// set the collection state ctor
	s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	return s.RowSourceImpl.Init(ctx, configData, connectionData, opts...)
}

func (s *AllColumnsSource) Identifier() string {
	return AllColumnsSourceIdentifier
}

func (s *AllColumnsSource) Collect(ctx context.Context) error {

	// populate enrichment fields the source is aware of
	// - in this case the connection
	sourceName := AllColumnsSourceIdentifier
	sourceEnrichmentFields := &schema.SourceEnrichment{
		CommonFields: schema.CommonFields{
			TpSourceName: &sourceName,
			TpSourceType: AllColumnsSourceIdentifier, // TODO: review this
		},
	}

	slog.Debug(">> Collecting data from source")

	for i := 1; i <= s.Config.RowCount; i++ {
		// populate the row data
		rowData := s.populateRowData(i)

		row := &types.RowData{Data: rowData, SourceEnrichment: sourceEnrichmentFields}
		slog.Debug(">> Sending row to plugin", row)
		if err := s.OnRow(ctx, row, nil); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}
	}

	return nil
}

func (s *AllColumnsSource) populateRowData(i int) *rows.AllColumns {
	return &rows.AllColumns{
		Id:        i,
		SmallInt:  int16(rand.Intn(32767)),
		Integer:   int32(rand.Int31()),
		BigInt:    rand.Int63(),
		UTinyInt:  uint8(rand.Intn(255)),
		UInteger:  rand.Uint32(),
		UBigInt:   rand.Uint64(),
		Float:     rand.Float32(),
		Double:    rand.Float64(),
		Decimal:   rand.Float64() * 100,
		Varchar:   randomString(10),
		Char:      string([]byte{byte(rand.Intn(26) + 65)}),
		Boolean:   rand.Intn(2) == 0,
		Date:      time.Now().AddDate(0, 0, rand.Intn(365)),
		Timestamp: time.Now().Add(time.Duration(rand.Intn(86400)) * time.Second),
		Interval:  fmt.Sprintf("%dh", rand.Intn(24)),
		IntArray:  []int32{rand.Int31(), rand.Int31(), rand.Int31()},
		CreatedAt: time.Now(),
	}
}

func randomString(n int) string {
	// Generates a random string of length n
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
