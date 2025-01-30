package chaos

import (
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-chaos/sources/all_columns"
	"github.com/turbot/tailpipe-plugin-chaos/sources/date_time"
	"github.com/turbot/tailpipe-plugin-chaos/sources/struct_columns"
	ac "github.com/turbot/tailpipe-plugin-chaos/tables/all_columns"
	dt "github.com/turbot/tailpipe-plugin-chaos/tables/date_time"
	sc "github.com/turbot/tailpipe-plugin-chaos/tables/struct_columns"
	"github.com/turbot/tailpipe-plugin-sdk/plugin"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

type Plugin struct {
	plugin.PluginImpl
}

func init() {
	// Register tables, with type parameters:
	// 1. row struct
	// 2. table implementation
	table.RegisterTable[*rows.AllColumns, *ac.AllColumnsTable]()
	table.RegisterTable[*rows.DateTime, *dt.DateTimeTable]()
	table.RegisterTable[*rows.StructColumns, *sc.StructColumnsTable]()

	// register sources
	row_source.RegisterRowSource[*all_columns.AllColumnsSource]()
	row_source.RegisterRowSource[*date_time.DateTimeSource]()
	row_source.RegisterRowSource[*struct_columns.StructColumnsSource]()
}

func NewPlugin() (_ plugin.TailpipePlugin, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = helpers.ToError(r)
		}
	}()

	p := &Plugin{
		PluginImpl: plugin.NewPluginImpl(config.PluginName),
	}

	return p, nil
}
