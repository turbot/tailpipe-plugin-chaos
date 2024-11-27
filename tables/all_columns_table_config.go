package tables

type AllColumnsTableConfig struct {
}

func (c *AllColumnsTableConfig) Validate() error {
	return nil
}

func (c *AllColumnsTableConfig) Identifier() string {
	return AllColumnsTableIdentifier
}
