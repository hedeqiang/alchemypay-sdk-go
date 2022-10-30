package services

import "github.com/hedeqiang/alchemypay-sdk-go/sdk"

const (
	defaultEndpoint = "127.0.0.1:9090"
)

type Client struct {
	sdk.Client
}

func NewClient(config *sdk.Config) (client *Client, err error) {
	client = &Client{}
	if config == nil {
		config = sdk.NewConfig().WithEndpoint(defaultEndpoint)
	}

	client.Init().WithConfig(config)
	return
}
