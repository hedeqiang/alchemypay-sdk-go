package sdk

import (
	"github.com/hedeqiang/alchemypay/sdk/response"
	"net/http"
)

func (c *Client) Notify(rawResponse *http.Response, resp response.Response) error {
	return response.ParseFromHttpResponse(rawResponse, resp)
}
