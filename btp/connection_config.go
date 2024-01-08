package btp

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type CISServiceKeyConfig struct {
	Endpoints map[string]string `json:"endpoints"`
	Uaa       map[string]string `json:"uaa"`
}

type BTPConfig struct {
	CISAccountServiceUrl      *string `hcl:"cis_accounts_service_url"`
	CISEntitlementsServiceUrl *string `hcl:"cis_entitlements_service_url"`
	AccessToken               *string `hcl:"cis_access_token"`
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

	// Reads the JSON file in cis_service_key_path and sets the value of cis_client_id and cis_client_secret
	if config.CISServiceKeyPath != nil {
		jsonFile, err := os.Open(*config.CISServiceKeyPath)

		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		data, err := io.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
		}

		var serviceKey CISServiceKeyConfig
		json.Unmarshal(data, &serviceKey)

		config.CISEventsServiceUrl = serviceKey.Endpoints["events"]

	}
	return config
}
