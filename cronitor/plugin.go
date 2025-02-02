package cronitor

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	plugin := &plugin.Plugin{
		Name: "steampipe-plugin-cronitor",
	}
	return plugin
}
