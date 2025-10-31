package brightness

import (
	"reflect"
	"testing"

	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/models"
)

func TestNewBrightnessPanel(t *testing.T) {
	light := models.Light{
		EntityID: "1", State: light_state.On, FriendlyName: "Living Room Light", Brightness: 75,
	}

	panel := InitializeNewBrightnessPanel(light)

	if panel == nil {
		t.Fatalf("InitializeNewBrightnessPanel() returned nil")
	}

	if !reflect.DeepEqual(&light, panel.light) {
		t.Errorf("InitializeNewBrightnessPanel() = %v, want %v", panel, light)
	}

	if panel.keys == nil {
		t.Errorf("InitializeNewBrightnessPanel() keys not initialized")
	}
}
