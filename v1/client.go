package bybit

import "time"

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
	}
}
