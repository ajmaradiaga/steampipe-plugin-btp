package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type BTPConfig struct {
	// Endpoints
	CISAccountServiceUrl      *string `cty:"cis_accounts_service_url"`
	CISEntitlementsServiceUrl *string `cty:"cis_entitlements_service_url"`

	// Username
	AccessToken *string `cty:"cis_access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"cis_accounts_service_url": {
		Type: schema.TypeString,
	},

	"cis_entitlements_service_url": {
		Type: schema.TypeString,
	},

	"cis_access_token": {
		Type:     schema.TypeString,
		Required: false,
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
