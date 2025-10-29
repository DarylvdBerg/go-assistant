package light_state

import (
	"errors"
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
	return stateName[ls]
}

func EnumValue(s string) (State, error) {
	for k, v := range stateName {
		if v == s {
			return k, nil
		}
	}

	return Unavailable, errors.New("unable to transform string to State enum entry")
}
