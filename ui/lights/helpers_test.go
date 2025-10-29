package lights

import (
	"reflect"
	"testing"

	"github.com/DarylvdBerg/go-assistant/shared/models"
)

func TestGetSelectedLight(t *testing.T) {
	lights := []models.Light{
		{"1", "on", "Light", 100},
		{"2", "off", "Light2", 50},
	}

	overview := InitLightOverview(lights)
	overview.list.Select(0)
	expected := &lights[0]

	selectedLight := overview.getSelectedLight()
	if selectedLight == nil || !reflect.DeepEqual(selectedLight, expected) {
		t.Errorf("GetSelectedLight() returned %+v, expected %+v", selectedLight, expected)
	}
}

func TestUpdateLightState(t *testing.T) {
	lights := []models.Light{
		{
			EntityID:     "1",
			State:        "on",
			FriendlyName: "Light",
			Brightness:   100,
		},
	}

	overview := InitLightOverview(lights)
	overview.list.Select(0)
	selectedLight := overview.getSelectedLight()
	if selectedLight == nil {
		t.Fatalf("GetSelectedLight() for initial light returned nil, expected a light")
	}

	if selectedLight.State != "on" {
		t.Errorf("Initial light state is %s, expected 'on'", selectedLight.State)
	}

	selectedLight.State = "off"
	overview.updateLightState(selectedLight)
	updatedLight := overview.getSelectedLight()
	if updatedLight == nil {
		t.Fatalf("GetSelectedLight() for updated light returned nil, expected a light")
	}

	if updatedLight.State != "off" {
		t.Errorf("Updated light state is %s, expected 'off'", updatedLight.State)
	}
}

func TestInitializeLightList(t *testing.T) {
	lights := []models.Light{
		{"1", "on", "Light", 100},
	}

	list := initializeLightList(lights)

	if len(list.Items()) != len(lights) {
		t.Errorf("InitializeLightList() returned %d items, expected %d", len(list.Items()), len(lights))
	}

	if !reflect.DeepEqual(list.Items()[0], lights[0]) {
		t.Errorf("InitializeLightList() returned item %+v, expected %+v", list.Items()[0], lights[0])
	}
}
