package btp

import (
	"context"
	"errors"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*Client, error) {
	// You can set custom *http.Client here
	client, err := NewClient(nil)
	if err != nil {
		return nil, err
	}

	accessToken := os.Getenv("BTP_ACCESS_TOKEN")

	btpConfig := GetConfig(d.Connection)
	if btpConfig.AccessToken != nil {
		accessToken = *btpConfig.AccessToken
	} else if accessToken != "" {
		btpConfig.AccessToken = &accessToken
	}

	if accessToken == "" {
		return nil, errors.New("'access_token' must be set in the connection configuration. Edit your connection configuration file or set the BTP_ACCESS_TOKEN environment variable and then restart Steampipe")
	}

	client.SetHeader("Authorization", "Bearer "+accessToken)

	// // example.zendesk.com
	// err = client.SetSubdomain(subdomain)
	// if err != nil {
	// 	return nil, err
	// }

	// // Authenticate with API token
	// client.SetCredential(zendesk.NewAPITokenCredential(user, token))

	return client, nil
}
