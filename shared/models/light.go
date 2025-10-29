package models

type Light struct {
	EntityID     string `json:"entity_id"`
	State        string `json:"state"` // TODO: State should eventually be an ENUM internally
	FriendlyName string `json:"friendly_name"`
	Brightness   int
}

func (l Light) Title() string       { return l.FriendlyName }
func (l Light) Description() string { return "" }
func (l Light) FilterValue() string { return l.FriendlyName }
