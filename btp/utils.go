package btp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type BTPService string

const (
	AccountsService    BTPService = "AccountsService"
	EntitlementService BTPService = "EntitlementService"
)

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

var defaultQueryStrings = map[string]string{}

type (
	// BTPClient of SAP BTP API
	BTPClient struct {
		httpClient *http.Client
		config     BTPConfig
		headers    map[string]string
		query      map[string]string
	}

	// BaseAPI encapsulates base methods for BTP client
	BaseAPI interface {
		Get(ctx context.Context, path string) ([]byte, error)
	}
)

// NewBTPClient creates new SAP BTP API client
func NewBTPClient(httpClient *http.Client, ctx context.Context, d *plugin.QueryData) (*BTPClient, error) {

	fnName := "NewBTPClient"
	debugFormat := "BTPClient.NewBTPClient: %s"
	logger := plugin.Logger(ctx)

	logger.Debug(fmt.Sprintf(debugFormat, "EXPLORING NewBTPClient\n===================="))

	// Load connection from cache
	cacheKey := d.Connection.Name
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*BTPClient), nil
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &BTPClient{httpClient: httpClient}
	connection := d.Connection

	if connection == nil || connection.Config == nil {
		client.config = BTPConfig{}
	} else {
		config, _ := connection.Config.(BTPConfig)

		logger.Debug(fmt.Sprintf(debugFormat, "connection is not nil"))

		logger.Debug(fmt.Sprintf(debugFormat, config.CISServiceKeyPath))

		CISServiceKeyPath := prioritiseEnvVar(os.Getenv("BTP_CIS_SERVICE_KEY_PATH"), config.CISServiceKeyPath)

		// Reads the JSON file in cis_service_key_path or the environment variable and sets the value of cis_client_id and cis_client_secret

		if CISServiceKeyPath != "" {

			logger.Debug(fmt.Sprintf(debugFormat, "CISServiceKeyPath is not empty"))

			expandedPath, err := homedir.Expand(*config.CISServiceKeyPath)
			if err != nil {
				return nil, err
			}

			jsonFile, err := os.Open(expandedPath)

			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			data, err := io.ReadAll(jsonFile)
			if err != nil {
				fmt.Println(err)
			}

			logger.Debug(fmt.Sprintf(debugFormat, string(data)))

			var serviceKey CISServiceKeyConfig
			err = json.Unmarshal(data, &serviceKey)

			logger.Warn(fnName, "data", data)
			logger.Warn(fnName, "err", err)

			if err != nil {
				logger.Error(fnName, "service_key_error", err)
				return nil, err
			}

			// Setting the config values from the service key
			config.CISAccountServiceUrl = prioritiseConfigVar(&config.CISAccountServiceUrl, serviceKey.Endpoints["accounts_service_url"])
			config.CISEntitlementsServiceUrl = prioritiseConfigVar(&config.CISEntitlementsServiceUrl, serviceKey.Endpoints["entitlements_service_url"])
			config.CISEventsServiceUrl = prioritiseConfigVar(&config.CISEventsServiceUrl, serviceKey.Endpoints["events_service_url"])
			config.CISTokenUrl = prioritiseConfigVar(&config.CISTokenUrl, serviceKey.Uaa["url"])
			config.CISClientId = prioritiseConfigVar(&config.CISClientId, serviceKey.Uaa["clientid"])
			config.CISClientSecret = prioritiseConfigVar(&config.CISClientSecret, serviceKey.Uaa["clientsecret"])

			logger.Debug(fmt.Sprintf(debugFormat, config.CISAccountServiceUrl))

		}

		client.config = config
	}

	client.headers = defaultHeaders
	client.query = defaultQueryStrings

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

// Retrieves the service URL from the BTP Config
func (b *BTPClient) getServiceURL(service BTPService) (string, error) {
	switch service {
	case AccountsService:
		return prioritiseEnvVar(os.Getenv("BTP_CIS_ACCOUNTS_SERVICE_URL"), &b.config.CISAccountServiceUrl), nil
	case EntitlementService:
		return prioritiseEnvVar(os.Getenv("BTP_CIS_ENTITLEMENTS_SERVICE_URL"), &b.config.CISEntitlementsServiceUrl), nil
	}

	return "", nil
}

// prepare request sets common request variables and sets headers and query strings
func (b *BTPClient) prepareRequest(ctx context.Context, service BTPService, req *http.Request, headers map[string]string, queryStrings map[string]string) (*http.Request, error) {
	out := req.WithContext(ctx)

	err := b.handleAuthentication(ctx, service)

	logger := plugin.Logger(ctx)

	if err != nil {
		logger.Error("BTPClient.prepareRequest", "connection_error", err)
		return nil, err
	}

	for key, value := range headers {
		b.headers[key] = value
	}

	for key, value := range queryStrings {
		b.query[key] = value
	}

	b.includeHeaders(out)
	b.includeQueryStrings(out)

	return out, nil
}

// includeHeaders set HTTP headers from client.headers to *http.Request
func (b *BTPClient) includeHeaders(req *http.Request) {
	for key, value := range b.headers {
		req.Header.Set(key, value)
	}
}

// includeHeaders set HTTP headers from client.query to *http.Request
func (b *BTPClient) includeQueryStrings(req *http.Request) {
	q := req.URL.Query() // Get a copy of the query values.

	for key, value := range b.query {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
}

func prioritiseConfigVar(configVar *string, serviceKeyValue string) string {
	// Prioritise config variable than service key value
	if configVar != nil && *configVar != "" {
		return *configVar
	}

	if serviceKeyValue != "" {
		return serviceKeyValue
	}

	return ""
}

func prioritiseEnvVar(envVar string, configVar *string) string {
	// Prioritise environment variable than config value
	if envVar != "" {
		return envVar
	}

	if configVar != nil {
		return *configVar
	}

	return ""
}

func (b *BTPClient) handleAuthentication(ctx context.Context, service BTPService) error {
	debugFormat := "BTPClient.handleAuthentication: %s"
	logger := plugin.Logger(ctx)

	logger.Debug(fmt.Sprintf(debugFormat, "EXPLORING b.config\n===================="))
	logger.Debug(fmt.Sprintf(debugFormat, *b.config.CISServiceKeyPath))
	logger.Debug(fmt.Sprintf(debugFormat, b.config.CISEventsServiceUrl))
	logger.Debug(fmt.Sprintf(debugFormat, "===================\nEND EXPLORING b.config\n===================="))

	logger.Debug(fmt.Sprintf(debugFormat, "service"), service)
	if service == AccountsService || service == EntitlementService {

		btpEnvironmentAccessToken := prioritiseEnvVar(os.Getenv("BTP_CIS_ACCESS_TOKEN"), &b.config.CISAccessToken)

		if btpEnvironmentAccessToken != "" {
			logger.Debug(fmt.Sprintf(debugFormat, "access token exists"))
			b.headers["Authorization"] = "Bearer " + btpEnvironmentAccessToken

			return nil
		} else {
			logger.Debug(fmt.Sprintf(debugFormat, "No access token, retrieving using service key details"))

			// Get access token using CIS service key details (client_id, client_secret, token_url)
			username := prioritiseEnvVar(os.Getenv("BTP_USERNAME"), b.config.Username)
			password := prioritiseEnvVar(os.Getenv("BTP_PASSWORD"), b.config.Password)
			tokenUrl := prioritiseEnvVar(os.Getenv("BTP_CIS_TOKEN_URL"), &b.config.CISTokenUrl)
			clientId := prioritiseEnvVar(os.Getenv("BTP_CIS_CLIENT_ID"), &b.config.CISClientId)
			clientSecret := prioritiseEnvVar(os.Getenv("BTP_CIS_CLIENT_SECRET"), &b.config.CISClientSecret)

			if username == "" || password == "" || tokenUrl == "" || clientId == "" || clientSecret == "" {
				err := errors.New("as no cis_access_token has been provided, you need to set 'username, password, cis_token_url, cis_client_id, cis_client_secret' in the connection configuration or as environment variables (BTP_USERNAME, BTP_PASSWORD, BTP_CIS_TOKEN_URL, BTP_CIS_CLIENT_ID, BTP_CIS_CLIENT_SECRET)")
				logger.Error("BTPClient.handleAuthentication", "connection_config_error", err)

				return err
			}

			// Check if tokenUrl contains /oauth/token, if not append it
			if !strings.Contains(tokenUrl, "/oauth/token") {
				logger.Debug(fmt.Sprintf(debugFormat, "Appending /oauth/token to tokenUrl"))
				tokenUrl = tokenUrl + "/oauth/token"
			}

			// Get access token using the details provided
			tokenResponse, err := getOauthAccessToken(ctx, tokenUrl, clientId, clientSecret, username, password)

			if err != nil {
				logger.Error("BTPClient.handleAuthentication", "authentication_error", err)
				return err
			}

			b.config.CISAccessToken = tokenResponse.AccessToken

			b.headers["Authorization"] = "Bearer " + tokenResponse.AccessToken

			return nil
		}

	} else {
		err := fmt.Errorf("service authentication not handled: %s.", service)
		logger.Error("BTPClient.handleAuthentication", "configuration_error", err)

		return err
	}

}

func getOauthAccessToken(ctx context.Context, tokenUrl string, clientId string, clientSecret string, username string, password string) (*TokenResponse, error) {
	logger := plugin.Logger(ctx)

	logger.Info("BTPClient.getOauthAccessToken", "Getting an access token")
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("grant_type=password&username=%s&password=%s", html.EscapeString(username), html.EscapeString(password)))

	client := &http.Client{}
	req, err := http.NewRequest(method, tokenUrl, payload)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(clientId, clientSecret)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse

	// Convert JSON response to Directory structure
	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

// Get from API
func (b *BTPClient) Get(ctx context.Context, service BTPService, path string, headers map[string]string, queryStrings map[string]string) ([]byte, error) {
	logger := plugin.Logger(ctx)
	logger.Debug("BTPClient.CONFIG", "ServiceKeyPath", *b.config.CISServiceKeyPath)
	logger.Debug("BTPClient.CONFIG", "EventsURL", b.config.CISEventsServiceUrl)
	baseURL, err := b.getServiceURL(service)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		logger.Error("BTPClient.Get", "connection_error", err)
		return nil, err
	}

	req, err = b.prepareRequest(ctx, service, req, headers, queryStrings)

	if err != nil {
		logger.Error("BTPClient.Get", "request_error", err)
		return nil, err
	}

	logger.Warn("BTPClient.Get", "requestURI", req.URL.String())

	resp, err := b.httpClient.Do(req)
	if err != nil {
		logger.Error("BTPClient.Get", "connection_error", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("BTPClient.Get", "connection_error", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("BTPClient.Get", "api_error", err)
		return nil, Error{
			body: body,
			resp: resp,
		}
	}
	return body, nil
}
