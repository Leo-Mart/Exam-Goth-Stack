package models

type KeystoneProfile struct {
	CurrentRating struct {
		Rating float64 `json:"rating"`
		Color  struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
	} `json:"current_mythic_rating"`
	CurrentPeriod struct {
		BestRuns []struct {
			Duration        int  `json:"duration"`
			KeystoneLevel   int  `json:"keystone_level"`
			CompletedInTime bool `json:"is_completed_within_time"`
			Affixes         []struct {
				AffixName string `json:"name"`
				AffixId   int    `json:"id"`
			} `json:"keystone_affixes"`
			Dungeon struct {
				DungeonName string `json:"name"`
			} `json:"dungeon"`
			Rating struct {
				Color struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"color"`
				Rating float64 `json:"rating"`
			} `json:"mythic_rating"`
		} `json:"best_runs"`
	} `json:"current_period"`
}
