package errors

import (
	"fmt"
)

type AlchemyPayError struct {
	Code string
	Msg  string
	Data interface{}
}

func (err *AlchemyPayError) Error() string {
	return fmt.Sprintf("[AlchemyPayError] code=%d, message=%s", err.Code, err.Msg)
}

func NewAlchemySDKError(code string, message string) error {
	return &AlchemyPayError{
		Code: code,
		Msg:  message,
	}
}

func (err *AlchemyPayError) GetCode() string {
	return err.Code
}

func (err *AlchemyPayError) GetMessage() string {
	return err.Msg
}

func (err *AlchemyPayError) GetData() interface{} {
	return err.Data
}
