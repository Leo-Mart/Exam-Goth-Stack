package models

type CharacterMedia struct {
	Character struct {
		Name string `json:"name"`
	} `json:"character"`
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}
