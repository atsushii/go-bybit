package bybit

import "time"

// Endpoints
const (
	baseAPIMainURL    = "https://api.bybit.com"
	baseAPITestnetURL = "https://api-testnet.bybit.com"
)

func currentTimestamp() int64 {
	return formatTimestamp(time.Now())
}

func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}


// Client define API client
type client struct {
	APIKey    string
	SecretKey string
	Debug     bool
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
