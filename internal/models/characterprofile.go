package models

type CharacterProfile struct {
	Name    string `json:"name"`
	Faction struct {
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Name string `json:"name"`
	} `json:"race"`
	Class struct {
		Name string `json:"name"`
	} `json:"character_class"`
	Spec struct {
		Name string `json:"name"`
	} `json:"active_spec"`
	Level              int `json:"level"`
	ItemLevelEquipped  int `json:"equippped_item_level"`
	KeystoneProfileURL struct {
		URL string `json:"href"`
	} `json:"mythic_keystone_profile"`
	EquipmentURL struct {
		URL string `json:"href"`
	} `json:"equipment"`
	MediaURL struct {
		URL string `json:"href"`
	} `json:"media"`
}
