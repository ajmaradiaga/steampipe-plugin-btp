package main

import (
	"github.com/ajmaradiaga/steampipe-plugin-btp/btp"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: btp.Plugin})
}
