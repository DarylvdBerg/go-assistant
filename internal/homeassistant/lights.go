package homeassistant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Light struct {
	EntityID 	string	`json:"entity_id"`
	State 		string	`json:"state"`
	FriendlyName	string `json:"friendly_name"`
}

func (c *Client) ListLights() ([]Light, error) {
	resp, err := c.doRequest("GET", "/api/states", nil);

	// If we received an error we shall return it.
	if err != nil {
		return nil, err
	}

	// Defer to close till the method completed execution.
	defer resp.Body.Close();
	
	// Since we cannot call specifically entities for lights, we'll have to filter.
	var allEntities []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&allEntities); err != nil {
		return nil, err;
	}
	
	var lights []Light
	for _, entity := range allEntities {
		if id, ok := entity["entity_id"].(string); ok && len(id) > 6 && id[:6] == "light." {
			light := Light {
				EntityID: id,
				State: entity["state"].(string),
			}
			if attrs, ok := entity["attributes"].(map[string]any); ok {
				if name, ok := attrs["friendly_name"].(string); ok {
					light.FriendlyName = name
				}
			}
			lights = append(lights, light)
		}
	}

	return lights, nil;
}

func (c *Client) TurnOnLight(entityID string) error {
	return c.callLightAction("turn_on", entityID)
}

func (c *Client) TurnOffLight(entityID string) error {
	return c.callLightAction("turn_off", entityID)
}

// Calls the Home assistant API and executes an action available for lights
func (c *Client) callLightAction(action, entityID string) error {
	path := fmt.Sprintf("/api/services/light/%s", action)
	body := map[string]any {
		"entity_id": entityID,
	}

	resp, err := c.doRequest("POST", path, body);
	if err != nil {
		return err;
	}

	defer resp.Body.Close();
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Home Assistant error: %s", string(b))
	}
	return nil;
}