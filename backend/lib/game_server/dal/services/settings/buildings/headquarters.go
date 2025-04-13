package buildings

import (
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
)

type HeadquartersLevel struct {
	Faction         enum.Faction
	Level           int
	Wood            int
	Stone           int
	Metal           int
	Population      int
	BuildSpeedMulti int
	TimeSeconds     int
}

type HeadquartersLevelDto struct {
	Faction         string `json:"faction"`
	Level           int    `json:"level"`
	Wood            int    `json:"wood"`
	Stone           int    `json:"stone"`
	Metal           int    `json:"metal"`
	Population      int    `json:"population"`
	BuildSpeedMulti int    `json:"build_speed_multi_x1000"`
	TimeSeconds     int    `json:"time_seconds"`
}

func (origin HeadquartersLevel) ToDto() (destination HeadquartersLevelDto) {
	destination.Faction = origin.Faction.String()
	destination.Level = origin.Level
	destination.Wood = origin.Wood
	destination.Stone = origin.Stone
	destination.Metal = origin.Metal
	destination.Population = origin.Population
	destination.BuildSpeedMulti = origin.BuildSpeedMulti
	destination.TimeSeconds = origin.TimeSeconds
	return destination
}

func (origin HeadquartersLevelDto) ToObj() (destination HeadquartersLevel, err error) {
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
	destination.BuildSpeedMulti = origin.BuildSpeedMulti
	destination.TimeSeconds = origin.TimeSeconds
	return destination, err
}

func (l *HeadquartersLevel) Sync(ctx context.Context, pool database.Querier) error {
	_, err := pool.Exec(
		ctx,
		`
		INSERT INTO settings_headquarters_levels (faction, level, wood, stone, metal, population, build_speed_multi_x1000, time_seconds) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (faction, level) DO UPDATE SET
			wood = EXCLUDED.wood,
			stone = EXCLUDED.stone,
			metal = EXCLUDED.metal,
			population = EXCLUDED.population,
			build_speed_multi_x1000 = EXCLUDED.build_speed_multi_x1000,
			time_seconds = EXCLUDED.time_seconds;
		`,
		l.Faction.String(), l.Level, l.Wood, l.Stone, l.Metal, l.Population, l.BuildSpeedMulti, l.TimeSeconds,
	)
	return err
}
