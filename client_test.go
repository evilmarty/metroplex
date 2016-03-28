package metroplex

import (
	"reflect"
	"testing"
)

func TestClient_Authorize(t *testing.T) {
	key := "X-Plex-Token"
	expected := "abc123"
	c := newClient().Authorize(expected)

	if actual := c.Headers[key]; actual != expected {
		t.Errorf("Expected the header %s to be %s, not %s.", key, expected, actual)
	}
}

func TestClient_Request(t *testing.T) {
	method, url := "POST", "https://www.example.com"
	r := newClient().Request(method, url)

	if actual := r.Request.Method; actual != method {
		t.Errorf("Expected the request method to be %s, not %s", method, actual)
	}

	if actual := r.Request.URL; actual.String() != url {
		t.Errorf("Expected the request url to be %s, not %s", url, actual)
	}

	for key, val := range DefaultHeaders {
		if r.Request.Header.Get(key) != val {
			t.Errorf("Expected the request to include the header %s: %s", key, val)
		}
	}
}

func TestNewClient(t *testing.T) {
	c := newClient()
	if !reflect.DeepEqual(c.Headers, DefaultHeaders) {
		t.Error("Expected the client to have default headers set")
	}
}
