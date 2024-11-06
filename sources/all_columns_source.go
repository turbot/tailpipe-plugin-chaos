package sources

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/parse"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/types"
	"golang.org/x/exp/rand"
)

const AllColumnsSourceIdentifier = "chaos"

// AllColumnsSource source is responsible for collecting audit logs from Turbot Pipes API
type AllColumnsSource struct {
	row_source.RowSourceImpl[*AllColumnsConfig]
}

func NewAllColumnsSource() row_source.RowSource {
	return &AllColumnsSource{}
}

func (s *AllColumnsSource) Init(ctx context.Context, configData *types.ConfigData, opts ...row_source.RowSourceOption) error {
	// // set the collection state ctor
	// s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	return s.RowSourceImpl.Init(ctx, configData, opts...)
}

func (s *AllColumnsSource) Identifier() string {
	return AllColumnsSourceIdentifier
}

func (s *AllColumnsSource) GetConfigSchema() parse.Config {
	return &AllColumnsConfig{}
}

func (s *AllColumnsSource) Collect(ctx context.Context) error {

	// populate enrichment fields the source is aware of
	// - in this case the connection
	sourceEnrichmentFields := &enrichment.CommonFields{
		TpSourceName: AllColumnsSourceIdentifier,
		TpSourceType: AllColumnsSourceIdentifier, // TODO: review this
	}

	slog.Debug(">> Collecting data from source")

	for i := 0; i < 10; i++ {
		rowData := &rows.AllColumns{
			SmallInt:  int16(rand.Intn(32767)),
			Integer:   int32(rand.Int31()),
			BigInt:    rand.Int63(),
			UTinyInt:  uint8(rand.Intn(255)),
			UInteger:  rand.Uint32(),
			UBigInt:   rand.Uint64(),
			Float:     rand.Float32(),
			Double:    rand.Float64(),
			Decimal:   rand.Float64() * 100, // Example range for decimal
			Varchar:   randomString(10),
			Char:      string([]byte{byte(rand.Intn(26) + 65)}), // Single ASCII character
			Boolean:   rand.Intn(2) == 0,
			Date:      time.Now().AddDate(0, 0, rand.Intn(365)),
			Timestamp: time.Now().Add(time.Duration(rand.Intn(86400)) * time.Second),
			Interval:  fmt.Sprintf("%dh", rand.Intn(24)), // Example interval as a string
			IntArray:  []int32{rand.Int31(), rand.Int31(), rand.Int31()},
			CreatedAt: time.Now(),
		}

		row := &types.RowData{Data: rowData, Metadata: sourceEnrichmentFields}
		slog.Debug(">> Sending row to plugin", row)
		if err := s.OnRow(ctx, row, nil); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}
	}

	return nil
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
