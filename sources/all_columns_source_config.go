package sources

import (
	"github.com/hashicorp/hcl/v2"
)

type AllColumnsConfig struct {
	// required to allow partial decoding
	Remain hcl.Body `hcl:",remain" json:"-"`
}

func (c *AllColumnsConfig) Validate() error {
	return nil
}
