package metroplex

import (
	"net/http"
)

const (
	SignInUrl = "https://my.plexapp.com/users/sign_in.xml"
)

type plexAuthQuery struct {
	AuthToken string `xml:"authenticationToken,attr"`
}

type Plex struct {
	client *Client
}

func SignIn(username, password string) (*Plex, error) {
	return signIn(username, password, SignInUrl)
}

func signIn(username, password, url string) (*Plex, error) {
	var q plexAuthQuery
	err := newClient().Request("POST", url).BasicAuth(username, password).RespondWith(http.StatusCreated).Xml(&q).Error
	if err != nil {
		return nil, err
	}

	return Auth(q.AuthToken)
}

func Auth(token string) (*Plex, error) {
	client := newClient().Authorize(token)
	return &Plex{client}, nil
}
