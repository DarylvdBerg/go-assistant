package homeassistant

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/DarylvdBerg/go-assistant/internal/mappers"
	"github.com/DarylvdBerg/go-assistant/shared/models"
)

const (
	ListLightsPath      = "/api/states"
	LightActionPath     = "/api/services/light/%s"
	LightBrightnessPath = "/api/services/light/turn_on"
)

// ListLights retrieves all entities starting with "light." and maps them to our Light model.
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
		light := mappers.MapToLight(entity)
		if light == nil {
			continue
		}
		lights = append(lights, *light)
	}

	return lights, nil
}

// ToggleLightState toggles the state of a light entity based on the provided action ("turn_on" or "turn_off").
func (hc *HaClient) ToggleLightState(entityID string, action string) error {
	path := fmt.Sprintf(LightActionPath, action)
	body := map[string]any{
		"entity_id": entityID,
	}

	return hc.callAction(path, body)
}

// ChangeBrightness changes the brightness of a light entity to the specified value (0-100).
func (hc *HaClient) ChangeBrightness(entityID string, brightness uint8) error {
	// The brightness value in home assistant is 255 for 100% and 2.5 for 1%, hence why we do the calculation.
	rightHand := float32(brightness) / 100
	brightnessValue := rightHand * 255

	body := map[string]any{
		"entity_id":  entityID,
		"brightness": brightnessValue,
	}

	return hc.callAction(LightBrightnessPath, body)
}

// ChangeColorTemp changes the color temperature of a light entity to the specified value in kelvin.
func (hc *HaClient) ChangeColorTemp(entityID string, colorTemp float64) error {
	body := map[string]any{
		"entity_id":  entityID,
		"color_temp": colorTemp,
	}

	return hc.callAction(LightBrightnessPath, body)
}
