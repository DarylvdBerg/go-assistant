package lights

import (
	"testing"

	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/models"
)

func TestInitLightOverview(t *testing.T) {
	lights := []models.Light{
		{"1", light_state.On, "Light", 100, nil, nil},
	}

	overview := InitLightOverview(lights)

	if len(overview.list.Items()) != 1 {
		t.Errorf("Expected 1 light in the list, got %d", len(overview.list.Items()))
	}

	if overview.list.Title == "" {
		t.Errorf("Expected a light title, got empty")
	}

	if overview.keys == nil {
		t.Errorf("Expected key bindings to be initialized, got nil")
	}
}
