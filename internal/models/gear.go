package models

type CharacterGear struct {
	Character struct {
		CharacterName string `json:"name"`
	} `json:"character"`
	EquippedItems []struct {
		Item struct {
			ItemId int `json:"id"`
		} `json:"item"`
		ItemName string `json:"name"`
	} `json:"equipped_items"`
}
