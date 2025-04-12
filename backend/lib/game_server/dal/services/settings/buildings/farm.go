package buildings

import (
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
)

type FarmLevel struct {
	Faction           enum.Faction
	Level             int
	Wood              int
	Stone             int
	Metal             int
	Population        int
	MaximumPopulation int
	TimeSeconds       int
}

type FarmLevelDto struct {
	Faction           string `json:"faction"`
	Level             int    `json:"level"`
	Wood              int    `json:"wood"`
	Stone             int    `json:"stone"`
	Metal             int    `json:"metal"`
	Population        int    `json:"population"`
	MaximumPopulation int    `json:"maximum_population"`
	TimeSeconds       int    `json:"time_seconds"`
}

func (origin FarmLevel) ToDto() (destination FarmLevelDto) {
	destination.Faction = origin.Faction.String()
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.MaximumPopulation = origin.MaximumPopulation
	destination.TimeSeconds = origin.TimeSeconds
	return destination
}

func (origin FarmLevelDto) ToObj() (destination FarmLevel, err error) {
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
	destination.MaximumPopulation = origin.MaximumPopulation
	destination.TimeSeconds = origin.TimeSeconds
	return destination, err
}

func (l *FarmLevel) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_farm_levels (faction, level, wood, stone, metal, population, maximum_population, time_seconds) 
		VALUES ($1, $2, $3, $4, $5, $,6, $7, $8)
		ON CONFLICT (faction, level) DO UPDATE SET
			wood = EXCLUDED.wood,
			stone = EXCLUDED.stone,
			metal = EXCLUDED.metal,
			population = EXCLUDED.population,
			maximum_population = EXCLUDED.maximum_population,
			time_seconds = EXCLUDED.time_seconds;
		`,
		l.Faction.String(), l.Level, l.Wood, l.Stone, l.Metal, l.Population, l.MaximumPopulation, l.TimeSeconds,
	)
	return err
}
