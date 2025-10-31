package supported_color_modes

import "github.com/DarylvdBerg/go-assistant/shared/utils"

type SupportedColorModes int

const (
	// None No color modes are available, only brightness
	None SupportedColorModes = iota
	// ColorTemp Temperature control is supported
	ColorTemp
)

var supportedColorModesName = map[SupportedColorModes]string{
	None:      "none",
	ColorTemp: "color_temp",
}

func (s SupportedColorModes) StringValue() string {
	return utils.StringValue(s, supportedColorModesName)
}

func EnumValue(s string) (SupportedColorModes, error) {
	return utils.EnumValue(s, supportedColorModesName, None)
}
