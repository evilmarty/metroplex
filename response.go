package metroplex

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	EmptyResponse = "HTTP/1.1 204 OK\n\n"
)

type Response struct {
	Error    error
	Response *http.Response
}

func (r *Response) Body(body *[]byte) *Response {
	if r.Error != nil {
		return r
	}

	res := r.Response
	b, err := ioutil.ReadAll(res.Body)
	*body = b
	return &Response{err, res}
}

func (r *Response) Xml(v interface{}) *Response {
	var body []byte
	if r.Body(&body).Error != nil {
		return r
	}

	err := xml.Unmarshal(body, &v)
	return &Response{err, r.Response}
}

func (r *Response) assertStatusCode(statusCodes ...int) *Response {
	if r.Error != nil {
		return r
	}

	statusCode := r.Response.StatusCode

	for _, sc := range statusCodes {
		if statusCode == sc {
			return r
		}
	}

	err := fmt.Errorf("Request received an unexpected status code of %d", statusCode)
	return &Response{err, r.Response}
}

func newResponse(r *Request) *Response {
	var res *http.Response
	var err error

	if r.Error != nil {
		buf := bufio.NewReader(strings.NewReader(EmptyResponse))
		res, _ = http.ReadResponse(buf, r.Request)
		err = r.Error
	} else {
		client := r.client
		res, err = client.Do(r.Request)
	}

	return &Response{err, res}
}
