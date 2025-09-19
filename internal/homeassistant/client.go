package homeassistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseUrl string
	Token string
	HTTPClient *http.Client
}

func CreateNewClient(baseUrl, token string) *Client {
	return &Client {
		BaseUrl: baseUrl,
		Token: token,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Builds up the necessary basics to do a request.
func (c *Client) doRequest(method, path string, body any) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseUrl, path);
	var buf bytes.Buffer

	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, &buf);
	if err != nil {
		return nil, err;
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	
	return c.HTTPClient.Do(req);
}
