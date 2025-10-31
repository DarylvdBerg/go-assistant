package mappers

import (
	"testing"

	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
)

func TestMapSupportedColorModes_NoSupportedModes_EmptyList(t *testing.T) {
	entries := []any{
		"invalid_mode",
		"another_invalid_mode",
	}

	modes := MapSupportedColorModes(entries)

	if len(modes) != 0 {
		t.Errorf("Expected 0 supported color modes, got %d", len(modes))
	}
}

func TestMapSupportedColorModes_SupportedModes_Contains_Brightness(t *testing.T) {
	entries := []any{
		"brightness",
	}

	modes := MapSupportedColorModes(entries)

	if len(modes) != 1 {
		t.Errorf("Expected 1 supported color mode, got %d", len(modes))
	}

	if modes[0] != supported_color_modes.Brightness {
		t.Errorf("Expected supported color mode to be Brightness, got %v", modes[0])
	}
}

func TestMapBrightness_FloatValue_ReturnsPercentage(t *testing.T) {
	brightness := 128.0

	percentage := MapBrightness(brightness)
	expected := int(brightness / 255 * 100)

	if percentage != expected {
		t.Errorf("Expected percentage to be %d, got %d", expected, percentage)
	}
}

func TestMapBrightness_IntValue_ReturnsPercentage(t *testing.T) {
	brightness := 64

	percentage := MapBrightness(brightness)
	expected := int(float64(brightness) / 255 * 100)

	if percentage != expected {
		t.Errorf("Expected percentage to be %d, got %d", expected, percentage)
	}
}

func TestMapBrightness_StringValue_ReturnsZero(t *testing.T) {
	brightness := "78"

	percentage := MapBrightness(brightness)
	if percentage != 0 {
		t.Errorf("Expected percentage to be 0 for invalid type, got %d", percentage)
	}

}
