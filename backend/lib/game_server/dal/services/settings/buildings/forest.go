package buildings

import (
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
)

type ForestLevel struct {
	Faction     enum.Faction
	Level       int
	Wood        int
	Stone       int
	Metal       int
	Population  int
	WoodHour    int
	TimeSeconds int
}

type ForestLevelDto struct {
	Faction     string `json:"faction"`
	Level       int    `json:"level"`
	Wood        int    `json:"wood"`
	Stone       int    `json:"stone"`
	Metal       int    `json:"metal"`
	Population  int    `json:"population"`
	WoodHour    int    `json:"wood_hour"`
	TimeSeconds int    `json:"time_seconds"`
}

func (origin ForestLevel) ToDto() (destination ForestLevelDto) {
	destination.Faction = origin.Faction.String()
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.WoodHour = origin.WoodHour
	destination.TimeSeconds = origin.TimeSeconds
	return destination
}

func (origin ForestLevelDto) ToObj() (destination ForestLevel, err error) {
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
	destination.WoodHour = origin.WoodHour
	destination.TimeSeconds = origin.TimeSeconds
	return destination, err
}

func (l *ForestLevel) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_forest_levels (faction, level, wood, stone, metal, population, wood_hour, time_seconds) 
		VALUES ($1, $2, $3, $4, $5, $,6, $7, $8)
		ON CONFLICT (faction, level) DO UPDATE SET
			wood = EXCLUDED.wood,
			stone = EXCLUDED.stone,
			metal = EXCLUDED.metal,
			population = EXCLUDED.population,
			wood_hour = EXCLUDED.wood_hour,
			time_seconds = EXCLUDED.time_seconds;
		`,
		l.Faction.String(), l.Level, l.Wood, l.Stone, l.Metal, l.Population, l.WoodHour, l.TimeSeconds,
	)
	return err
}
