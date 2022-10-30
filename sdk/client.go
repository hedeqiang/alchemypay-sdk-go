package sdk

import (
	"alchemypay/sdk/request"
	"alchemypay/sdk/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Client is the base struct of service clients
type Client struct {
	signMethod string
	Config     *Config
}

func (c *Client) Init() *Client {
	c.signMethod = "MD5"
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	return c
}

// Send send the request and return the response to the client.
// Parameter request accepts concrete request object which follow Request.
func (c *Client) Send(req request.Request, resp response.Response) error {
	method := req.GetMethod()
	builder := GetParameterBuilder(method)

	params := req.GetParams()
	if _, ok := params["merchantCode"]; !ok {
		req.SetParams("merchantCode", c.Config.MerchantCode)
	}
	if len(params) > 1 {
		signStr := OrderParam(params, c.Config.PrivateKey)
		req.SetParams("sign", SignWithMd5(signStr))
	}

	jsonReq, _ := json.Marshal(req)

	endPoint := c.Config.Endpoint

	reqUrl := fmt.Sprintf("%s://%s/%s/%s", c.Config.Scheme, endPoint, req.GetURL(), req.GetAction())

	body, err := builder.BuildBody(jsonReq)
	if err != nil {
		return err
	}

	rawResponse, err := c.doSend(method, reqUrl, body, req.GetHeaders())
	if err != nil {
		return err
	}

	return response.ParseFromHttpResponse(rawResponse, resp)
}

func (c *Client) doSend(method, url, data string, header map[string]string) (*http.Response, error) {
	client := &http.Client{Timeout: c.Config.Timeout}

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	c.setHeader(req, header)

	return client.Do(req)
}

func (c *Client) setHeader(req *http.Request, header map[string]string) {
	req.Header.Set("Content-Type", "application/json")

	for k, v := range header {
		req.Header.Set(k, v)
	}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
