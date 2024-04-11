package go_http_client

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func NewRestyClient(clientConfig *ClientConfig) *resty.Client {
	c := resty.New()
	if clientConfig != nil {
		if clientConfig.Endpoint != "" {
			c.SetBaseURL(clientConfig.Endpoint)
		}
		if clientConfig.Timeout > 0 {
			c.SetTimeout(time.Duration(clientConfig.Timeout) * time.Second)
		}
		if clientConfig.Debug {
			c.SetDebug(true)
		}
		if clientConfig.Auth != nil && clientConfig.Auth.Scheme != "" && clientConfig.Auth.Token != "" {
			c.SetAuthScheme(clientConfig.Auth.Scheme)
			c.SetAuthToken(clientConfig.Auth.Token)
		}
	}
	return c
}
