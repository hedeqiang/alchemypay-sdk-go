package services

import (
	"encoding/json"
	"github.com/hedeqiang/alchemypay/sdk/request"
	"github.com/hedeqiang/alchemypay/sdk/response"
)

type AlchemyRequest struct {
	*request.BaseRequest
	Params map[string]interface{}
}

func (r *AlchemyRequest) GetParams() map[string]interface{} {
	if r.Params == nil {
		return make(map[string]interface{})
	}
	return r.Params
}

func (r *AlchemyRequest) SetParams(key string, value interface{}) {
	if r.Params == nil {
		r.Params = make(map[string]interface{})
	}
	r.Params[key] = value
}

func NewPayRequest() (req *AlchemyRequest) {
	req = &AlchemyRequest{
		BaseRequest: &request.BaseRequest{
			Method: "POST",
		},
	}
	return
}

type AlchemyResponse struct {
	*response.BaseResponse
}

func NewPayResponse() *AlchemyResponse {
	return &AlchemyResponse{
		BaseResponse: &response.BaseResponse{},
	}
}

func (r *AlchemyResponse) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r *AlchemyResponse) GetResponseBody() string {
	data, _ := json.Marshal(r.Data)
	return string(data)
}

func (c *Client) Pay(req *AlchemyRequest) (resp *AlchemyResponse, err error) {
	if req == nil {
		req = NewPayRequest()
	}

	resp = NewPayResponse()
	err = c.Send(req, resp)
	return
}
