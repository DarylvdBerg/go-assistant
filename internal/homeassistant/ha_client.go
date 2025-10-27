package homeassistant

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/DarylvdBerg/go-assistant/internal/client"
)

type HaClient struct {
	*client.Client
}

func (hc *HaClient) callAction(path string, body map[string]any) error {
	resp, err := hc.Request("POST", path, body)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}

	}(resp.Body)
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
