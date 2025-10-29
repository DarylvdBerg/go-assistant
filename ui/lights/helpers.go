package lights

import (
	"log"

	"github.com/DarylvdBerg/go-assistant/internal/homeassistant"
	"github.com/DarylvdBerg/go-assistant/shared"
	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/ui/style"
	"github.com/charmbracelet/bubbles/list"
)

func (e LightList) getSelectedLight() *models.Light {
	selectedItem := e.list.SelectedItem()
	if selectedItem == nil {
		return nil
	}

	light, ok := selectedItem.(models.Light)
	if !ok {
		return nil
	}

	return &light
}

func toggleLight(light *models.Light) {
	var action string
	if light.State == shared.LightStateOn {
		action = TurnOffAction
		light.State = shared.LightStateOff
	} else {
		action = TurnOnAction
		light.State = shared.LightStateOn
	}

	err := homeassistant.GetClient().ToggleLightState(light.EntityID, action)
	if err != nil {
		log.Fatal("failed to toggle light state: ", err)
	}
}

func (e LightList) updateLightState(updatedLight *models.Light) {
	items := e.list.Items()
	for i, item := range items {
		if light, ok := item.(models.Light); ok {
			if light.EntityID == updatedLight.EntityID {
				// Update the light's state
				light.State = updatedLight.State
				light.Brightness = updatedLight.Brightness
				// Replace the item in the list
				e.list.SetItem(i, light)
				break
			}
		}
	}
}

func initializeLightList(lights []models.Light) list.Model {
	items := make([]list.Item, 0)

	for _, s := range lights {
		items = append(items, s)
	}

	return list.New(items, style.NewOverviewStyleDelegate(), 0, 0)
}
