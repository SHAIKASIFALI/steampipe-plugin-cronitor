package main

import (
	"github.com/SHAIKASIFALI/steampipe-plugin-cronitor/cronitor"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: cronitor.Plugin,
	})
}
