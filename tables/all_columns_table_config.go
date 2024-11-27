package tables

type AllColumnsTableConfig struct {
	Index string `json:"index" hcl:"index"`
}

func (c *AllColumnsTableConfig) Validate() error {
	return nil
}

func (c *AllColumnsTableConfig) Identifier() string {
	return AllColumnsTableIdentifier
}
