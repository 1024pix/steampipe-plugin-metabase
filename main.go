package main

import (
	"github.com/1024pix/steampipe-plugin-metabase/metabase"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: metabase.Plugin})
}
