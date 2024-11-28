package sources

import (
	"github.com/hashicorp/hcl/v2"
)

type StructColumnsSourceConfig struct {
	// required to allow partial decoding
	Remain hcl.Body `hcl:",remain" json:"-"`
	// required to set the row count
	RowCount int `hcl:"row_count" json:"row_count"`
}

func (c *StructColumnsSourceConfig) Validate() error {
	// if nothing is explicitly set, we'll default to 10 rows
	if c.RowCount == 0 {
		c.RowCount = 10
	}
	return nil
}

func (c *StructColumnsSourceConfig) Identifier() string {
	return StructColumnsSourceIdentifier
}
