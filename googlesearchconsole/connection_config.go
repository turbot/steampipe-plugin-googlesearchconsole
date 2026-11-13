package googlesearchconsole

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/schema"
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
	if connection == nil || connection.GetConfig() == nil {
		return gscConfig{}
	}
	config, _ := connection.GetConfig().(gscConfig)
	return config
}
