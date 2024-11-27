package sources

import (
	"github.com/hashicorp/hcl/v2"
)

type AllColumnsSourceConfig struct {
	// required to allow partial decoding
	Remain hcl.Body `hcl:",remain" json:"-"`
	// required to set the row count
	RowCount int `hcl:"row_count" json:"row_count"`
}

func (c *AllColumnsSourceConfig) Validate() error {
	// if nothing is explicitly set, we'll default to 10 rows
	if c.RowCount == 0 {
		c.RowCount = 10
	}
	return nil
}

func (c *AllColumnsSourceConfig) Identifier() string {
	return AllColumnsSourceIdentifier
}
