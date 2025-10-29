package brightness

import (
	"reflect"
	"testing"

	"github.com/DarylvdBerg/go-assistant/shared/models"
)

func TestNewBrightnessPanel(t *testing.T) {
	light := models.Light{
		EntityID: "1", State: "on", FriendlyName: "Living Room Light", Brightness: 75,
	}

	panel := NewBrightnessPanel(light)

	if panel == nil {
		t.Fatalf("NewBrightnessPanel() returned nil")
	}

	if !reflect.DeepEqual(&light, panel.light) {
		t.Errorf("NewBrightnessPanel() = %v, want %v", panel, light)
	}

	if panel.keys == nil {
		t.Errorf("NewBrightnessPanel() keys not initialized")
	}
}
