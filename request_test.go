package metroplex

import (
	"net/http"
	"testing"
)

func TestRequest_Headers(t *testing.T) {
	cases := map[string]string{
		"X-Foo": "foo",
		"X-Bar": "bar",
	}
	r := newRequest("GET", "https://www.example.com")
	r = r.Headers(cases)

	for c, expected := range cases {
		if r.Request.Header.Get(c) != expected {
			t.Errorf("Expected the request header %s: %s", c, expected)
		}
	}
}

func TestRequest_BasicAuth(t *testing.T) {
	r := newRequest("GET", "https://www.example.com")
	r = r.BasicAuth("chuck", "nope")

	if r.Request.Header.Get("Authorization") == "" {
		t.Errorf("Expected the authorization header to be set")
	}
}

func TestRequest_RespondWith(t *testing.T) {
	r := newRequest("GET", "https://www.example.com").RespondWith(http.StatusOK)
	if r.Error != nil {
		t.Error(r.Error)
	}
}
