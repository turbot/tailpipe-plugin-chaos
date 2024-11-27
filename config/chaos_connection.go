package config

const PluginName = "chaos"

type ChaosConnection struct {
}

func (c *ChaosConnection) Validate() error {
	return nil
}

func (c *ChaosConnection) Identifier() string {
	return PluginName
}
