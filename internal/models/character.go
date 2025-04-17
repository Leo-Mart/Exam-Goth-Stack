package models

import "time"

type RaiderioCharacter struct {
	Name       string `json:"name"`
	Race       string `json:"race"`
	Class      string `json:"class"`
	Spec       string `json:"active_spec_name"`
	Thumbnail  string `json:"thumbnail_url"`
	RecentRuns []struct {
		Dungeon             string    `json:"dungeon"`
		ShortName           string    `json:"short_name"`
		MythicLevel         int       `json:"mythic_level"`
		CompletedAt         time.Time `json:"completed_at"`
		ClearTimeMs         int       `json:"clear_time_ms"`
		ParTimeMs           int       `json:"par_time_ms"`
		NumKeystoneUpgrades int       `json:"num_keystone_upgrades"`
		Score               float64   `json:"score"`
		Affixes             []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Icon string `json:"icon"`
		} `json:"affixes"`
		URL string `json:"url"`
	} `json:"mythic_plus_recent_runs"`
}

type BlizzardCharacter struct {
	ID               string
	CharacterProfile BlizzCharProfile
	KeystoneProfile  KeystoneProfile
	Gear             CharacterGear
}
