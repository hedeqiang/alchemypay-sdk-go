package services

import (
	"encoding/json"
	"github.com/hedeqiang/alchemy/sdk/request"
	"github.com/hedeqiang/alchemy/sdk/response"
)

type AlchemyRequest struct {
	*request.BaseRequest
	Params map[string]string
}

func (r *AlchemyRequest) GetParams() map[string]string {
	if r.Params == nil {
		return make(map[string]string)
	}
	return r.Params
}

func (r *AlchemyRequest) SetParams(key, value string) {
	if r.Params == nil {
		r.Params = make(map[string]string)
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
