package models

import (
	"errors"
	"strings"
)

type BuildingName int

const (
	Headquarters BuildingName = iota
	Farm
	Warehouse
	Forest
	Quarry
	Mine
)

type BuildingLevel struct {
	Building    BuildingName
	Faction     Faction
	Wood        int
	Stone       int
	Metal       int
	Population  int
	TimeSeconds int
}

func (d BuildingName) String() string {
	return [...]string{"headquarters", "farm", "warehouse", "forest", "quarry", "mine"}[d]
}

func BuildingNameFromString(str string) (BuildingName, error) {
	switch strings.ToLower(str) {
	case "headquarters":
		return Headquarters, nil
	case "farm":
		return Farm, nil
	case "warehouse":
		return Warehouse, nil
	case "forest":
		return Forest, nil
	case "quary":
		return Quarry, nil
	case "mine":
		return Mine, nil
	default:
		return 0, errors.New("invalid faction: " + str)
	}
}
