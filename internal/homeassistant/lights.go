package homeassistant

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/DarylvdBerg/go-assistant/shared/models"
)

const (
	ListLightsPath      = "/api/states"
	LightActionPath     = "/api/services/light/%s"
	LightBrightnessPath = "/api/services/light/turn_on"
)

const (
	LightsPart = "light."
)

func (hc *HaClient) ListLights() ([]models.Light, error) {
	res, err := hc.Request("GET", ListLightsPath, nil)

	// If we received an error we shall return it.
	if err != nil {
		return nil, err
	}

	// Defer to close till the method completed execution.
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	// Since we cannot call specifically entities for lights, we'll have to filter.
	var allEntities []map[string]any
	if err := json.NewDecoder(res.Body).Decode(&allEntities); err != nil {
		return nil, err
	}

	lights := make([]models.Light, 0)
	for _, entity := range allEntities {
		light := mapToLight(entity)
		if light == nil {
			continue
		}
		lights = append(lights, *light)
	}

	return lights, nil
}

func mapToLight(entity map[string]any) *models.Light {
	id, ok := entity["entity_id"].(string)
	if !ok || len(id) <= 6 || id[:6] != "light." {
		return nil
	}

	state, ok := entity["state"].(string)
	if !ok {
		return nil
	}

	attrs, ok := entity["attributes"].(map[string]any)
	if !ok {
		return nil
	}

	name, ok := attrs["friendly_name"].(string)
	if !ok {
		return nil
	}

	light := &models.Light{
		EntityID:     id,
		State:        state,
		FriendlyName: name,
	}

	if brightness, ok := attrs["brightness"]; ok {
		light.Brightness = haBrightnessToPercent(brightness)
	}

	return light
}

// Converts Home Assistant brightness (0–255) to percentage (0–100)
func haBrightnessToPercent(brightnessValue any) int {
	switch v := brightnessValue.(type) {
	case float64:
		return int(v / 255 * 100)
	case int:
		return int(float64(v) / 255 * 100)
	default:
		return 0
	}
}

func (ha *HaClient) ToggleLightState(entityID string, action string) error {
	path := fmt.Sprintf(LightActionPath, action)
	body := map[string]any{
		"entity_id": entityID,
	}

	return ha.callAction(path, body)
}

func (ha *HaClient) ChangeBrightness(entityID string, brightness uint8) error {
	// The brightness value in home assistant is 255 for 100% and 2.5 for 1%, hence why we do the calculation.
	rightHand := float32(brightness) / 100
	brightnessValue := rightHand * 255

	body := map[string]any{
		"entity_id":  entityID,
		"brightness": brightnessValue,
	}

	return ha.callAction(LightBrightnessPath, body)
}
