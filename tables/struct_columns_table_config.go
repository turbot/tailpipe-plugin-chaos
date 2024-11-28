package tables

type StructColumnsTableConfig struct {
}

func (c *StructColumnsTableConfig) Validate() error {
	return nil
}

func (c *StructColumnsTableConfig) Identifier() string {
	return StructColumnsTableIdentifier
}
