package main

import (
	"github.com/turbot/steampipe-plugin-googlesearchconsole/googlesearchconsole"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: googlesearchconsole.Plugin})
}
