package models

type Specialization struct {
	Specializations []struct {
		Specialization struct {
			Name string `json:"name"`
		} `json:"specialization"`
		Loadouts []struct {
			IsActive             bool   `json:"is_active"`
			LoadoutCode          string `json:"talent_loadout_code"`
			SelectedClassTalents []struct {
				Tooltip struct {
					SpellTooltip struct {
						Spell struct {
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
			} `json:"selected_class_talents"`
			SelectedSpecTalents []struct {
				Tooltip struct {
					SpellTooltip struct {
						Spell struct {
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
			} `json:"selected_spec_talents"`
			SelectedHeroTalents []struct {
				Tooltip struct {
					SpellTooltip struct {
						Spell struct {
							Name string `json:"name"`
							Id   int    `json:"id"`
						} `json:"spell"`
					} `json:"spell_tooltip"`
				} `json:"tooltip"`
			} `json:"selected_hero_talents"`
		} `json:"loadouts"`
	} `json:"specializations"`
}
