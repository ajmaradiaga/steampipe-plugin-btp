package btp

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

type (
	// Client of Zendesk API
	Client struct {
		baseURL    *url.URL
		httpClient *http.Client
		headers    map[string]string
	}

	// BaseAPI encapsulates base methods for zendesk client
	BaseAPI interface {
		Get(ctx context.Context, path string) ([]byte, error)
		Post(ctx context.Context, path string, data interface{}) ([]byte, error)
		Put(ctx context.Context, path string, data interface{}) ([]byte, error)
		Delete(ctx context.Context, path string) error
	}
)

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{httpClient: httpClient}
	client.headers = defaultHeaders
	return client, nil
}

// SetHeader saves HTTP header in client. It will be included all API request
func (z *Client) SetHeader(key string, value string) {
	z.headers[key] = value
}

// SetEndpointURL replace full URL of endpoint without subdomain validation.
// This is mainly used for testing to point to mock API server.
func (z *Client) SetEndpointURL(newURL string) error {
	baseURL, err := url.Parse(newURL)
	if err != nil {
		return err
	}

	z.baseURL = baseURL
	return nil
}

// get get JSON data from API and returns its body as []bytes
func (z *Client) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, z.baseURL.String()+path, nil)
	if err != nil {
		return nil, err
	}

	req = z.prepareRequest(ctx, req)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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

// prepare request sets common request variables such as authn and user agent
func (z *Client) prepareRequest(ctx context.Context, req *http.Request) *http.Request {
	out := req.WithContext(ctx)
	z.includeHeaders(out)
	// if z.credential != nil {
	// 	if z.credential.Bearer() {
	// 		out.Header.Add("Authorization", "Bearer "+z.credential.Secret())
	// 	} else {
	// 		out.SetBasicAuth(z.credential.Email(), z.credential.Secret())
	// 	}
	// }

	return out
}

// includeHeaders set HTTP headers from client.headers to *http.Request
func (z *Client) includeHeaders(req *http.Request) {
	for key, value := range z.headers {
		req.Header.Set(key, value)
	}
}

// addOptions build query string
// func addOptions(s string, opts interface{}) (string, error) {
// 	u, err := url.Parse(s)
// 	if err != nil {
// 		return s, err
// 	}

// 	qs, err := query.Values(opts)
// 	if err != nil {
// 		return s, err
// 	}

// 	u.RawQuery = qs.Encode()
// 	return u.String(), nil
// }

// Get allows users to send requests not yet implemented
func (z *Client) Get(ctx context.Context, path string) ([]byte, error) {
	return z.get(ctx, path)
}

// Post allows users to send requests not yet implemented
// func (z *Client) Post(ctx context.Context, path string, data interface{}) ([]byte, error) {
// 	return z.post(ctx, path, data)
// }

// // Put allows users to send requests not yet implemented
// func (z *Client) Put(ctx context.Context, path string, data interface{}) ([]byte, error) {
// 	return z.put(ctx, path, data)
// }

// // Delete allows users to send requests not yet implemented
// func (z *Client) Delete(ctx context.Context, path string) error {
// 	return z.delete(ctx, path)
// }
