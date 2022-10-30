package request

type Request interface {
	GetMethod() string
	GetURL() string
	GetAction() string
	GetHeaders() map[string]string
	GetParams() map[string]string
	SetParams(key, value string)
}

// BaseRequest Request is the base struct of service requests
type BaseRequest struct {
	Method string
	Header map[string]string
	URL    string
	Action string
}

func (r BaseRequest) GetMethod() string {
	return r.Method
}

func (r BaseRequest) GetHeaders() map[string]string {
	return r.Header
}

func (r BaseRequest) GetAction() string {
	return r.Action
}

func (r BaseRequest) GetURL() string {
	return r.URL
}

// AddHeader only adds pin or erp, they will be encoded to base64 code
func (r *BaseRequest) AddHeader(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}
