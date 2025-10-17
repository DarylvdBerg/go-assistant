package homeassistant

import (
	"fmt"
	"go-assistant/internal/client"
	"io"
	"net/http"
)

type HaClient struct {
	*client.Client
}

func (hc *HaClient) callAction(path string, body map[string]any) error {
	resp, err := hc.Request("POST", path, body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Home Assistant error: %s", string(b))
	}
	return nil
}

func GetClient() *HaClient {
	return &HaClient{
		Client: client.GetClient(),
	}
}
