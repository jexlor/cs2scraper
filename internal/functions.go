package internal

import "github.com/jexlor/cs2scraper/scraper/config"

func RemoveDuplicates(skins []config.Skin) []config.Skin {
	seen := make(map[string]bool)
	var unique []config.Skin

	for _, s := range skins {
		key := s.Name + "|" + s.Weapon + "|" + s.Rarity
		if !seen[key] {
			seen[key] = true
			unique = append(unique, s)
		}
	}
	return unique
}
