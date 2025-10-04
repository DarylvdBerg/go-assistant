package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Client struct {
	BaseUrl    string
	Token      string
	HTTPClient *http.Client
}

var client *Client

func setClient(c *Client) {
	client = c
}

func GetClient() *Client {
	return client
}

func NewClient() *Client {

	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	client := &Client{
		BaseUrl: config.BaseUrl,
		Token:   config.Token,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	setClient(client)

	return client
}

func (c *Client) Request(method, path string, body any) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseUrl, path)
	var buf bytes.Buffer

	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	return c.HTTPClient.Do(req)
}
