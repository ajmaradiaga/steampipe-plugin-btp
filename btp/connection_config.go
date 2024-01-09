package btp

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type CISServiceKeyConfig struct {
	Endpoints map[string]string `json:"endpoints"`
	Uaa       map[string]string `json:"uaa"`
}

type BTPConfig struct {
	Username                  *string `hcl:"username"`
	Password                  *string `hcl:"password"`
	CISServiceKeyPath         *string `hcl:"cis_service_key_path"`
	CISAccessToken            string  `hcl:"cis_access_token,optional"`
	CISAccountServiceUrl      string  `hcl:"cis_accounts_service_url,optional"`
	CISEntitlementsServiceUrl string  `hcl:"cis_entitlements_service_url,optional"`
	CISEventsServiceUrl       string  `hcl:"cis_events_service_url,optional"`
	CISTokenUrl               string  `hcl:"cis_token_url,optional"`
	CISClientId               string  `hcl:"cis_client_id,optional"`
	CISClientSecret           string  `hcl:"cis_client_secret,optional"`
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
