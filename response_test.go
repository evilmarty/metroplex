package metroplex

import (
	"net/http"
	"testing"
)

func TestResponse_Body(t *testing.T) {
	var actual []byte
	expected := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	s := serve(200, string(expected))
	defer s.Close()

	req := newRequest("GET", s.URL)
	res := newResponse(req)

	if err := res.Body(&actual).Error; err != nil {
		t.Fatalf("Did not expect to receive the error %s", err)
	}
	if string(actual) != string(expected) {
		t.Errorf("Expected the body to be %s, not %s", expected, actual)
	}
}

func TestResponse_Xml(t *testing.T) {
	type Person struct {
		Name  string `xml:"FullName"`
		Email string `xml:"email,attr"`
	}

	data := `
		<Person email="chuck@example.com">
			<FullName>Chuck Testa</FullName>
		</Person>
	`
	expected := Person{"Chuck Testa", "chuck@example.com"}
	actual := Person{}

	s := serve(200, data)
	defer s.Close()

	req := newRequest("GET", s.URL)
	res := newResponse(req)

	if err := res.Xml(&actual).Error; err != nil {
		t.Fatalf("Did not expect to receive the error %s", err)
	}
	if actual != expected {
		t.Errorf("Expected the body to be %s, not %s", expected, actual)
	}
}

func TestResponse_assertStatusCode(t *testing.T) {
	s := serve(404, "ooops")
	defer s.Close()

	req := newRequest("GET", s.URL)
	res := newResponse(req)

	if err := res.assertStatusCode(http.StatusNotFound).Error; err != nil {
		t.Fatalf("Did not expect to receive the error %s", err)
	}

	if res.assertStatusCode(http.StatusOK).Error == nil {
		t.Fatalf("Expected to receive an error but did not")
	}
}

func TestNewResponse_requestError(t *testing.T) {
	req := newRequest(" ", " ")
	if req.Error == nil {
		t.Fatal("Expected the request to contain an error", req)
	}

	res := newResponse(req)
	if res.Error != req.Error {
		t.Errorf("Expected the response to have the error %s, not %s", req.Error, res.Error)
	}
}

func TestNewResponse_requestGood(t *testing.T) {
	req := newRequest("", "https://www.example.com")
	if req.Error != nil {
		t.Fatal("Did not expect the request to contain the error %s", req.Error)
	}

	res := newResponse(req)
	if res.Error != nil {
		t.Errorf("Did not expect the response to contain the error", res.Error)
	}
}
