package metroplex

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {
	data := `
		<user email="chuck@example.com" id="1" uuid="abc123" thumb="https://plex.tv/users/abc123/avatar" username="chucktesta" title="chucktesta" locale="" authenticationToken="abc123" scrobbleTypes="" restricted="0" home="0" guest="0"  secure="1" certificateVersion="2">
		</user>
	`
	expected := "abc123"
	s := serve(201, data)
	defer s.Close()

	p, err := signIn("chuck", "nope", s.URL)
	if err != nil {
		t.Fatalf("Did not expect the error %s", err)
	}
	if actual := p.client.Headers["X-Plex-Token"]; actual != expected {
		t.Errorf("Expected the auth token to be %s, not %s", expected, actual)
	}
}

func TestAuth(t *testing.T) {
	expected := "abc123"
	p, err := Auth(expected)
	if err != nil {
		t.Fatalf("Did not expect the error %s", err)
	}
	if actual := p.client.Headers["X-Plex-Token"]; actual != expected {
		t.Errorf("Expected the auth token to be %s, not %s", expected, actual)
	}
}

func serve(status int, content string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "text/xml")
		w.Write([]byte(content))
	}))
}
