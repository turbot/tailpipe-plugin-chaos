package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
)

// DateTime is the struct containing the enriched data
type DateTime struct {
	// embed required enrichment fields
	schema.CommonFields

	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}
