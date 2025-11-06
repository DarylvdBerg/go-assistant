package light_state

import "testing"

func TestLightState_StringValue(t *testing.T) {
	state := On
	actual := state.StringValue()
	if actual != "on" {
		t.Errorf("Expected 'on', got '%s'", actual)
	}
}

func TestLightState_EnumValue(t *testing.T) {
	state := "off"
	actual := EnumValue(state)

	if actual != Off {
		t.Errorf("Expected Off, got %v", actual)
	}
}
