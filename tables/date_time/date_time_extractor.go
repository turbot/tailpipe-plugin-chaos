package date_time

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
)

// DateTimeExtractor is an extractor that receives JSON serialised DateTimeBatch objects
// and extracts DateTime records from them
type DateTimeExtractor struct {
}

// NewDateTimeExtractor creates a new DateTimeExtractor
func NewDateTimeExtractor() artifact_source.Extractor {
	return &DateTimeExtractor{}
}

func (c *DateTimeExtractor) Identifier() string {
	return "cloudtrail_log_extractor"
}

// Extract unmarshalls the artifact data as an DateTimeBatch and returns the DateTime records
func (c *DateTimeExtractor) Extract(_ context.Context, a any) ([]any, error) {
	// the expected input type is a JSON byte[] deserializable to DateTimeBatch
	jsonBytes, ok := a.([]byte)
	if !ok {
		return nil, fmt.Errorf("expected byte[], got %T", a)
	}

	// decode json ito DateTimeBatch
	var logs []rows.DateTime
	err := json.Unmarshal(jsonBytes, &logs)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	slog.Debug("DateTimeExtractor", "record count", len(logs))
	var res = make([]any, len(logs))
	for i, record := range logs {
		res[i] = &record
	}
	return res, nil
}
