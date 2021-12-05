package bybit

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type params map[string]interface{}

type request struct {
	method string
	endpoint string
	query url.Values
	form url.Values
	header http.Header
	body io.Reader
	fullURL string
	sign string
}

type RequestOption func(*request)

func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set params with key/values to request form body
func (r *request) setFormParams(m params) *request {
	for k, v := range m {
		r.setFormParam(k, v)
	}
	return r
}