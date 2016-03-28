package metroplex

var (
	DefaultHeaders = map[string]string{
		"X-Plex-Platform":          "golang",
		"X-Plex-Platform-Version":  "0.0",
		"X-Plex-Provides":          "player,controller",
		"X-Plex-Version":           "0.0",
		"X-Plex-Device":            "platform",
		"X-Plex-Client-Identifier": "metroplex",
	}
)

type Client struct {
	Headers map[string]string
}

func (c *Client) Authorize(token string) *Client {
	c.Headers["X-Plex-Token"] = token
	return c
}

func (c *Client) Request(method, url string) *Request {
	return newRequest(method, url).Headers(c.Headers)
}

func newClient() *Client {
	return &Client{
		Headers: DefaultHeaders,
	}
}
