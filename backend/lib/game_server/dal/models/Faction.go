package models

import (
	"errors"
	"strings"
)

type FactionEnum int

const (
	CALDARI FactionEnum = iota
	VARNAK
	DAWNHOLD
)

func (d FactionEnum) String() string {
	return [...]string{"caldari", "varnak", "dawnhold"}[d]
}

func FromString(factionStr string) (FactionEnum, error) {
	switch strings.ToLower(factionStr) {
	case "caldari":
		return CALDARI, nil
	case "varnak":
		return VARNAK, nil
	case "dawnhold":
		return DAWNHOLD, nil
	default:
		return 0, errors.New("invalid faction: " + factionStr)
	}
}
