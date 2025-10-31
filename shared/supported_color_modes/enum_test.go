package supported_color_modes

import "testing"

func TestSupportedColorModes_StringValue(t *testing.T) {
	mode := ColorTemp
	actual := mode.StringValue()

	if actual != "color_temp" {
		t.Errorf("Expected 'color_temp', got '%s'", actual)
	}
}

func TestSupportedColorModes_EnumValue(t *testing.T) {
	mode := "color_temp"
	actual := EnumValue(mode)

	if actual != ColorTemp {
		t.Errorf("Expected ColorTemp, got '%v'", actual)
	}
}
