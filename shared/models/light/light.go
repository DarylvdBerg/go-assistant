package models

type Light struct {
	EntityID     string `json:"entity_id"`
	State        string `json:"state"`
	FriendlyName string `json:"friendly_name"`
	Brightness   int    `json:"brightness"`
}

func (l Light) Title() string       { return l.FriendlyName }
func (l Light) Description() string { return "" }
func (l Light) FilterValue() string { return l.FriendlyName }
