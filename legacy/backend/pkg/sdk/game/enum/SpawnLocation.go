package enum

import (
	"errors"
	"strings"
)

type SpawnLocation int

const (
	Random SpawnLocation = iota
	SouthWest
	SouthEast
	NorthWest
	NorthEast
)

func (d SpawnLocation) String() string {
	return [...]string{"random", "south_west", "south_east", "north_west", "north_east"}[d]
}

func SpawnLocationFromString(factionStr string) (SpawnLocation, error) {
	switch strings.ToLower(factionStr) {
	case "random":
		return Random, nil
	case "south_west":
		return SouthWest, nil
	case "south_east":
		return SouthEast, nil
	case "north_west":
		return NorthWest, nil
	case "north_east":
		return NorthEast, nil
	default:
		return 0, errors.New("invalid faction: " + factionStr)
	}
}
