package models

import (
	"errors"
	"strings"
)

type RealmStatus int

const (
	OPEN RealmStatus = iota
	CLOSED
	ENDED
)

func (d RealmStatus) String() string {
	return [...]string{"open", "closed", "ended"}[d]
}

func FromString(statusStr string) (RealmStatus, error) {
	switch strings.ToLower(statusStr) {
	case "open":
		return OPEN, nil
	case "closed":
		return CLOSED, nil
	case "ended":
		return ENDED, nil
	default:
		return 0, errors.New("invalid status: " + statusStr)
	}
}
