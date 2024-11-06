package config

import "github.com/turbot/tailpipe-plugin-sdk/parse"

type ChaosConnection struct {
}

func NewChaosConnection() parse.Config {
	return &ChaosConnection{}
}

func (c *ChaosConnection) Validate() error {
	return nil
}
