package gsc

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type gscConfig struct {
	Credentials *string `cty:"credentials"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"credentials": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &gscConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) gscConfig {
	if connection == nil || connection.Config == nil {
		return gscConfig{}
	}
	config, _ := connection.Config.(gscConfig)
	return config
}
