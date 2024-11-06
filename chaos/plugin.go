package chaos

import (
	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/sources"
	"github.com/turbot/tailpipe-plugin-chaos/tables"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
	//"log/slog"
	//"time"
)

type Plugin struct {
	plugin.PluginImpl
}

func NewPlugin() (plugin.TailpipePlugin, error) {
	//slog.Info("Pipes Plugin starting")
	//time.Sleep(10 * time.Second)
	//slog.Info("Pipes Plugin started")

	p := &Plugin{
		PluginImpl: plugin.NewPluginImpl("chaos", config.NewChaosConnection),
	}

	// register the tables, sources and mappers that we provide
	resources := &plugin.ResourceFunctions{
		Tables:  []func() table.Table{tables.NewAllColumnsTable},
		Sources: []func() row_source.RowSource{sources.NewAllColumnsSource},
	}

	if err := p.RegisterResources(resources); err != nil {
		return nil, err
	}

	return p, nil
}
