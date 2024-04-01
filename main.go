package main

import (
	"github.com/turbot/steampipe-plugin-gsc/gsc"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: gsc.Plugin})
}
