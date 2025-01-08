package sources

import (
	"github.com/hashicorp/hcl/v2"
)

type UnicodeColumnsSourceConfig struct {
	Remain   hcl.Body `hcl:",remain" json:"-"`
	RowCount int      `hcl:"row_count" json:"row_count"`
}

func (c *UnicodeColumnsSourceConfig) Validate() error {
	if c.RowCount == 0 {
		c.RowCount = 1
	}
	return nil
}

func (c *UnicodeColumnsSourceConfig) Identifier() string {
	return UnicodeColumnsSourceIdentifier
}
