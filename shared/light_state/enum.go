package light_state

import (
	"github.com/DarylvdBerg/go-assistant/shared/utils"
)

type State int

const (
	On State = iota
	Off
	Unavailable
)

var stateName = map[State]string{
	On:          "on",
	Off:         "off",
	Unavailable: "unavailable",
}

func (ls State) StringValue() string {
	return utils.StringValue(ls, stateName)
}

func EnumValue(s string) (State, error) {
	return utils.EnumValue(s, stateName, Unavailable)
}
