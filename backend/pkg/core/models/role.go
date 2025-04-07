package models

import (
	"errors"
	"strings"
)

type Role int

const (
	ADMIN Role = iota
	MODERATOR
	NPC
	PLAYER
	GUEST
)

func (d Role) String() string {
	return [...]string{"admin", "moderator", "npc", "player", "guest"}[d]
}

func FromString(roleStr string) (Role, error) {
	switch strings.ToLower(roleStr) {
	case "admin":
		return ADMIN, nil
	case "moderator":
		return MODERATOR, nil
	case "npc":
		return NPC, nil
	case "player":
		return PLAYER, nil
	case "guest":
		return GUEST, nil
	default:
		return 0, errors.New("invalid role: " + roleStr)
	}
}
