package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
)

type UnicodeColumns struct {
	// embed required enrichment fields
	schema.CommonFields
	Title          string                  `json:"title"`
	Timestamp      time.Time               `json:"timestamp"`
	TopLevelString string                  `json:"top_level_string"`
	JsonObject     *map[string]interface{} `json:"json_object,omitempty"`
}
