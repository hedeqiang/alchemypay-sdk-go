package response

import (
	"alchemypay/sdk/errors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type ErrorResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

type BaseResponse struct {
	ErrorResponse
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) error {
	if err := json.Unmarshal(body, r); err != nil {
		return err
	}
	if r.Code == "" {
		return errors.NewAlchemySDKError(r.Code, r.Msg)
	}

	return nil
}

func ParseFromHttpResponse(rawResponse *http.Response, response Response) error {
	defer rawResponse.Body.Close()
	body, err := io.ReadAll(rawResponse.Body)
	if err != nil {
		return err
	}
	if rawResponse.StatusCode != 200 {
		return fmt.Errorf("request fail with status: %s, with body: %s", rawResponse.Status, body)
	}

	if err := response.ParseErrorFromHTTPResponse(body); err != nil {
		return err
	}

	return json.Unmarshal(body, &response)
}
