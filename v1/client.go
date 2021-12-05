package bybit

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Endpoints
const (
	baseAPIMainURL    = "https://api.bybit.com"
	baseAPITestnetURL = "https://api-testnet.bybit.com"
)

// useTestnet switch all the API endpoints from production to the testnet
var useTestnet = false

func currentTimestamp() int64 {
	return formatTimestamp(time.Now())
}

// time into Unix timestamp in milliseconds
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// Client define API client
type client struct {
	APIKey    string
	SecretKey string
	BaseURL    string
	Debug     bool
}

// return endpoint of the REST API according to the useTestnet flag
func getAPIEndpoint() string {
	if useTestnet {
		return baseAPITestnetURL
	}
	return baseAPIMainURL
}

// newClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func newClient(apiKey, secretKey string) *client {
	return &client{
		APIKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:    getAPIEndpoint(),
	}
}

func (c *client) parseRequest(r *request, opts ...RequestOption) (err error) {
	//set request options from user
	for _, opt := range opts {
		opt(r)
	}
	// TODO: add validator
	err = r.validate()
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	// TODO: from lone 235
}
func (c *client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
}
