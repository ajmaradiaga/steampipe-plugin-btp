package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type BTPConfig struct {
	// UAA
	UAAAPIUrl       *string `cty:"uaa_apiurl"`
	UAAClientId     *string `cty:"uaa_clientid"`
	UAAClientSecret *string `cty:"uaa_clientsecret"`

	// Endpoints
	CISAccountServiceUrl      *string `cty:"cis_accounts_service_url"`
	CISEntitlementsServiceUrl *string `cty:"cis_entitlements_service_url"`

	// Username
	Username    *string `cty:"username"`
	Password    *string `cty:"password"`
	AccessToken *string `cty:"cis_access_token"`
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

	"cis_accounts_service_url": {
		Type: schema.TypeString,
	},

	"cis_entitlements_service_url": {
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
