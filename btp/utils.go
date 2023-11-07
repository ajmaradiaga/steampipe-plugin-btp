package btp

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"

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

	// BaseAPI encapsulates base methods for zendesk client
	BaseAPI interface {
		Get(ctx context.Context, path string) ([]byte, error)
	}
)

// NewBTPClient creates new SAP BTP API client
func NewBTPClient(httpClient *http.Client, connection *plugin.Connection) (*BTPClient, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &BTPClient{httpClient: httpClient}

	if connection == nil || connection.Config == nil {
		client.config = BTPConfig{}
	} else {
		config, _ := connection.Config.(BTPConfig)
		client.config = config
	}

	client.headers = defaultHeaders
	client.query = defaultQueryStrings

	return client, nil
}

// Retrieves the service URL from the BTP Config
func (b *BTPClient) getServiceURL(service BTPService) (string, error) {
	switch service {
	case AccountsService:
		btpEnvironmentVariable := os.Getenv("BTP_CIS_ACCOUNTS_SERVICE_URL")

		// Prioritise environment variable
		if btpEnvironmentVariable != "" {
			return btpEnvironmentVariable, nil
		}

		return *b.config.CISAccountServiceUrl, nil
	case EntitlementService:
		btpEnvironmentVariable := os.Getenv("BTP_CIS_ENTITLEMENTS_SERVICE_URL")

		// Prioritise environment variable
		if btpEnvironmentVariable != "" {
			return btpEnvironmentVariable, nil
		}

		return *b.config.CISEntitlementsServiceUrl, nil
	}

	return "", nil
}

// prepare request sets common request variables and sets headers and query strings
func (b *BTPClient) prepareRequest(ctx context.Context, req *http.Request, headers map[string]string, queryStrings map[string]string) *http.Request {
	out := req.WithContext(ctx)

	b.handleAuthentication()

	for key, value := range headers {
		b.headers[key] = value
	}

	for key, value := range queryStrings {
		b.query[key] = value
	}

	b.includeHeaders(out)
	b.includeQueryStrings(out)

	return out
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

func (b *BTPClient) handleAuthentication() error {
	btpEnvironmentAccessToken := os.Getenv("BTP_CIS_ACCESS_TOKEN")

	// Prioritise access token set as an environment variable
	if btpEnvironmentAccessToken != "" {
		b.config.AccessToken = &btpEnvironmentAccessToken
	} else if b.config.AccessToken != nil {
		btpEnvironmentAccessToken = *b.config.AccessToken
	}

	if btpEnvironmentAccessToken == "" {
		return errors.New("'cis_access_token' must be set in the connection configuration. Edit your connection configuration file or set the BTP_CIS_ACCESS_TOKEN environment variable and then restart Steampipe")
	}

	b.headers["Authorization"] = "Bearer " + btpEnvironmentAccessToken

	return nil

}

// Get from API
func (b *BTPClient) Get(ctx context.Context, service BTPService, path string, headers map[string]string, queryStrings map[string]string) ([]byte, error) {
	logger := plugin.Logger(ctx)

	baseURL, err := b.getServiceURL(service)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	req = b.prepareRequest(ctx, req, headers, queryStrings)

	logger.Warn("Get", "requestURI", req.URL.String())

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}
	return body, nil
}
