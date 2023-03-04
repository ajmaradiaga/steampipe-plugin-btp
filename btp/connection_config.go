package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type BTPConfig struct {
	// UAA
	UAAAPIUrl       *string `cty:"uaa_apiurl"`
	UAAClientId     *string `cty:"uaa_clientid"`
	UAAClientSecret *string `cty:"uaa_clientsecret"`

	// Endpoints
	EndpointsAccountServiceUrl *string `cty:"endpoints_accounts_service_url"`

	// Username
	Username    *string `cty:"username"`
	Password    *string `cty:"password"`
	AccessToken *string `cty:"access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"uaa_apiurl": {
		Type:     schema.TypeString,
		Required: false,
	},

	"uaa_clientid": {
		Type:     schema.TypeString,
		Required: false,
	},

	"uaa_clientsecret": {
		Type:     schema.TypeString,
		Required: false,
	},

	"endpoints_accounts_service_url": {
		Type: schema.TypeString,
	},

	"username": {
		Type:     schema.TypeString,
		Required: false,
	},

	"password": {
		Type:     schema.TypeString,
		Required: false,
	},

	"access_token": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func ConfigInstance() interface{} {
	return &BTPConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) BTPConfig {
	if connection == nil || connection.Config == nil {
		return BTPConfig{}
	}
	config, _ := connection.Config.(BTPConfig)
	return config
}
