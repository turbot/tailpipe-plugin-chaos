package main

import (
	"log/slog"

	"github.com/turbot/tailpipe-plugin-chaos/chaos"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
)

func main() {
	err := plugin.Serve(&plugin.ServeOpts{
		PluginFunc: chaos.NewPlugin,
	})

	if err != nil {
		slog.Error("Error starting plugin", "error", err)
	}
}
