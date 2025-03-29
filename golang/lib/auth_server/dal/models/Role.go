package models

type Role int

const (
	ADMIN Role = iota
	MODERATOR
	NPC
	PLAYER
	GUEST
)

func (d Role) String() string {
	return [...]string{"ADMIN", "MODERATOR", "NPC", "PLAYER", "GUEST"}[d]
}
