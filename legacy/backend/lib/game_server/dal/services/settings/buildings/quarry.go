package buildings

import (
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
)

type QuarryLevel struct {
	Faction     enum.Faction
	Level       int
	Wood        int
	Stone       int
	Metal       int
	Population  int
	StoneHour   int
	Points      int
	TimeSeconds int
}

type QuarryLevelDto struct {
	Faction     string `json:"faction"`
	Level       int    `json:"level"`
	Wood        int    `json:"wood"`
	Stone       int    `json:"stone"`
	Metal       int    `json:"metal"`
	Population  int    `json:"population"`
	StoneHour   int    `json:"stone_hour"`
	Points      int    `json:"points"`
	TimeSeconds int    `json:"time_seconds"`
}

func (origin QuarryLevel) ToDto() (destination QuarryLevelDto) {
	destination.Faction = origin.Faction.String()
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.StoneHour = origin.StoneHour
	destination.Points = origin.Points
	destination.TimeSeconds = origin.TimeSeconds
	return destination
}

func (origin QuarryLevelDto) ToObj() (destination QuarryLevel, err error) {
	faction, err := enum.FactionFromString(origin.Faction)
	if err != nil {
		return destination, err
	}
	destination.Faction = faction
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.StoneHour = origin.StoneHour
	destination.Points = origin.Points
	destination.TimeSeconds = origin.TimeSeconds
	return destination, err
}

func (l *QuarryLevel) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_quarry_levels (faction, level, wood, stone, metal, population, stone_hour, points, time_seconds) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (faction, level) DO UPDATE SET
			wood = EXCLUDED.wood,
			stone = EXCLUDED.stone,
			metal = EXCLUDED.metal,
			population = EXCLUDED.population,
			stone_hour = EXCLUDED.stone_hour,
			points = EXCLUDED.points,
			time_seconds = EXCLUDED.time_seconds;
		`,
		l.Faction.String(), l.Level, l.Wood, l.Stone, l.Metal, l.Population, l.StoneHour, l.Points, l.TimeSeconds,
	)
	return err
}
