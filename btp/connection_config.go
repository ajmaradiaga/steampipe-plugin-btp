package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type btpConfig struct {
	// UAAAPIUrl       *string `cty:"uaa_apiurl"`
	// UAAClientId     *string `cty:"uaa_clientid"`
	// UAAClientSecret *string `cty:"uaa_clientsecret"`
	EndpointsAccountServiceUrl *string `cty:"endpoints_accounts_service_url"`
	Username                   *string `cty:"username"`
	Password                   *string `cty:"password"`
	AccessToken                *string
}

var ConfigSchema = map[string]*schema.Attribute{
	// "uaa_apiurl": {
	// 	Type: schema.TypeString,
	// },

	// "uaa_clientid": {
	// 	Type: schema.TypeString,
	// },

	// "uaa_clientsecret": {
	// 	Type: schema.TypeString,
	// },

	"endpoints_accounts_service_url": {
		Type: schema.TypeString,
	},

	"username": {
		Type: schema.TypeString,
	},

	"password": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &btpConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) btpConfig {
	if connection == nil || connection.Config == nil {
		return btpConfig{}
	}
	config, _ := connection.Config.(btpConfig)
	return config
}
