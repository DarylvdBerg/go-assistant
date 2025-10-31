package config

type Config struct {
	BaseUrl string `json:"home_assistant_endpoint"`
	Token   string `json:"token"`
}
