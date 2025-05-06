package models

type CharacterGear struct {
	EquippedItems []struct {
		Item struct {
			ItemId int `json:"id"`
		} `json:"item"`
		ItemName string `json:"name"`
	} `json:"equipped_items"`
}
