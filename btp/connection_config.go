package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type BTPConfig struct {
	CISAccountServiceUrl      *string `hcl:"cis_accounts_service_url"`
	CISEntitlementsServiceUrl *string `hcl:"cis_entitlements_service_url"`
	AccessToken               *string `hcl:"cis_access_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"cis_accounts_service_url": {
		Type: schema.TypeString,
	},

	"cis_entitlements_service_url": {
		Type: schema.TypeString,
	},

	"cis_access_token": {
		Type: schema.TypeString,
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
