package btp

// func connect(ctx context.Context, d *plugin.QueryData) (*BTPClient, error) {
// 	// You can set custom *http.Client here
// 	client, err := NewBTPClient(nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	btpEnvironmentAccessToken := os.Getenv("BTP_ACCESS_TOKEN")

// 	btpConfig := GetConfig(d.Connection)

// 	// Prioritise access token set as an environment variable
// 	if btpEnvironmentAccessToken != "" {
// 		btpConfig.AccessToken = &btpEnvironmentAccessToken
// 	} else if btpConfig.AccessToken != nil {
// 		btpEnvironmentAccessToken = *btpConfig.AccessToken
// 	}

// 	if btpEnvironmentAccessToken == "" {
// 		return nil, errors.New("'access_token' must be set in the connection configuration. Edit your connection configuration file or set the BTP_ACCESS_TOKEN environment variable and then restart Steampipe")
// 	}

// 	client.SetHeader("Authorization", "Bearer "+btpEnvironmentAccessToken)

// 	return client, nil
// }
