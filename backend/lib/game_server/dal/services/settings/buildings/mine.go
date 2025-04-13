package buildings

import (
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
)

type MineLevel struct {
	Faction     enum.Faction
	Level       int
	Wood        int
	Stone       int
	Metal       int
	Population  int
	MetalHour   int
	TimeSeconds int
}

type MineLevelDto struct {
	Faction     string `json:"faction"`
	Level       int    `json:"level"`
	Wood        int    `json:"wood"`
	Stone       int    `json:"stone"`
	Metal       int    `json:"metal"`
	Population  int    `json:"population"`
	MetalHour   int    `json:"metal_hour"`
	TimeSeconds int    `json:"time_seconds"`
}

func (origin MineLevel) ToDto() (destination MineLevelDto) {
	destination.Faction = origin.Faction.String()
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.MetalHour = origin.MetalHour
	destination.TimeSeconds = origin.TimeSeconds
	return destination
}

func (origin MineLevelDto) ToObj() (destination MineLevel, err error) {
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
	destination.MetalHour = origin.MetalHour
	destination.TimeSeconds = origin.TimeSeconds
	return destination, err
}

func (l *MineLevel) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_mine_levels (faction, level, wood, stone, metal, population, metal_hour, time_seconds) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (faction, level) DO UPDATE SET
			wood = EXCLUDED.wood,
			stone = EXCLUDED.stone,
			metal = EXCLUDED.metal,
			population = EXCLUDED.population,
			metal_hour = EXCLUDED.metal_hour,
			time_seconds = EXCLUDED.time_seconds;
		`,
		l.Faction.String(), l.Level, l.Wood, l.Stone, l.Metal, l.Population, l.MetalHour, l.TimeSeconds,
	)
	return err
}
