package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
)

// AllColumns is the struct containing the enriched data
type AllColumns struct {
	// embed required enrichment fields
	enrichment.CommonFields

	// Additional fields
	// Numeric types
	SmallInt int16   `json:"smallint_column"`
	Integer  int32   `json:"integer_column"`
	BigInt   int64   `json:"bigint_column"`
	UTinyInt uint8   `json:"utinyint_column"`
	UInteger uint32  `json:"uinteger_column"`
	UBigInt  uint64  `json:"ubigint_column"`
	Float    float32 `json:"float_column"`
	Double   float64 `json:"double_column"`
	Decimal  float64 `json:"decimal_column"`

	// String types
	Varchar string `json:"varchar_column"`
	Char    string `json:"char_column"`

	// Boolean type
	Boolean bool `json:"boolean_column"`

	// Date/Time types
	Date      time.Time `json:"date_column"`
	Timestamp time.Time `json:"timestamp_column"`

	// Interval type
	Interval string `json:"interval_column"` // Using string as placeholder for interval

	// Array type (simplified to slice of integers for this example)
	IntArray []int32 `json:"int_array_column"`

	// CreatedAt time.Time `json:"created_at"`
	CreatedAt time.Time `json:"created_at"`
}
