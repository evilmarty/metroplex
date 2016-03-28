package metroplex

import (
	"net/http"
)

type Request struct {
	Error   error
	Request *http.Request
	client  *http.Client
}

func (r *Request) Headers(headers map[string]string) *Request {
	if r.Error == nil {
		req := r.Request
		for name, val := range headers {
			req.Header.Set(name, val)
		}
	}

	return r
}

func (r *Request) BasicAuth(username, password string) *Request {
	if r.Error == nil {
		r.Request.SetBasicAuth(username, password)
	}

	return r
}

func (r *Request) RespondWith(statusCodes ...int) *Response {
	return newResponse(r).assertStatusCode(statusCodes...)
}

func newRequest(method, url string) *Request {
	req, err := http.NewRequest(method, url, nil)
	client := &http.Client{}
	return &Request{err, req, client}
}
