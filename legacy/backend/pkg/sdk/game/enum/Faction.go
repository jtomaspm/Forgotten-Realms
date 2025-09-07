package enum

import (
	"errors"
	"strings"
)

type Faction int

const (
	Caldari Faction = iota
	Varnak
	Dawnhold
	Forgotten
)

func (d Faction) String() string {
	return [...]string{"caldari", "varnak", "dawnhold", "forgotten"}[d]
}

func FactionFromString(factionStr string) (Faction, error) {
	switch strings.ToLower(factionStr) {
	case "caldari":
		return Caldari, nil
	case "varnak":
		return Varnak, nil
	case "dawnhold":
		return Dawnhold, nil
	case "forgotten":
		return Forgotten, nil
	default:
		return 0, errors.New("invalid faction: " + factionStr)
	}
}
