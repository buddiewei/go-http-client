package go_http_client

import (
	"context"
	"fmt"
	"testing"
)

func TestNewRestyClient(t *testing.T) {
	conf := ClientConfig{
		Endpoint: "http://www.baidu.com",
		Timeout:  2,
		Debug:    false,
		Auth:     nil,
	}
	client := NewRestyClient(&conf)
	response, err := client.R().SetContext(context.Background()).Get("/s?wd=golang")
	if response.StatusCode() != 200 {
		t.Fatalf("status code is not 200, but %d", response.StatusCode())
	}
	fmt.Println(response, err)
}
