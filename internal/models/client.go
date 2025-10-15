package models

import "net/http"

type Client struct {
	BaseUrl    string
	Token      string
	HTTPClient *http.Client
}
