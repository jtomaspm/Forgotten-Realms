package village_s

import (
	"backend/lib/game_server/dal/services/chunks"
	"backend/lib/game_server/dal/services/settings"
	"backend/lib/game_server/dal/services/villages"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/google/uuid"
)

type ChunkRelativePosition struct {
	CoordX         int
	CoordY         int
	DistanceCenter float32
}

type Vector struct {
	BaseX      int
	BaseY      int
	DirectionX int
	DirectionY int
}

func GetVector(location enum.SpawnLocation, realmSettings settings.Realm) (Vector, error) {
	quadrant_size := realmSettings.MapSize / 2
	switch location {
	case enum.NorthEast:
		return Vector{
			BaseX:      quadrant_size,
			BaseY:      quadrant_size - 1,
			DirectionX: 1,
			DirectionY: -1,
		}, nil
	case enum.NorthWest:
		return Vector{
			BaseX:      quadrant_size - 1,
			BaseY:      quadrant_size - 1,
			DirectionX: -1,
			DirectionY: -1,
		}, nil
	case enum.SouthEast:
		return Vector{
			BaseX:      quadrant_size,
			BaseY:      quadrant_size,
			DirectionX: 1,
			DirectionY: 1,
		}, nil
	case enum.SouthWest:
		return Vector{
			BaseX:      quadrant_size - 1,
			BaseY:      quadrant_size,
			DirectionX: -1,
			DirectionY: 1,
		}, nil
	case enum.Random:
		return GetVector(enum.SpawnLocation(rand.Intn(4)+1), realmSettings)
	default:
		return Vector{}, fmt.Errorf("invalid location: %s", location)
	}
}

func SpawnVillage(ctx context.Context, db database.Querier, playerId uuid.UUID, location enum.SpawnLocation, realmSettings settings.Realm) (villages.NewVillage, error) {
	areaSizeChunks := (realmSettings.MapSize / 2) / realmSettings.ChunkSize

	var chunkPositions []ChunkRelativePosition
	for x := range areaSizeChunks {
		for y := range areaSizeChunks {
			chunkPositions = append(chunkPositions, ChunkRelativePosition{
				CoordX:         x,
				CoordY:         y,
				DistanceCenter: float32(math.Sqrt(float64(x*x + y*y))),
			})
		}
	}
	sort.Slice(chunkPositions, func(i, j int) bool {
		return chunkPositions[i].DistanceCenter < chunkPositions[j].DistanceCenter
	})

	targetPopulation := float32(realmSettings.ChunkFillPercent / 100)
	if targetPopulation > 1 {
		return villages.NewVillage{}, fmt.Errorf("maximum population is 100%%, current is %d%%", realmSettings.ChunkFillPercent)
	}
	for _, position := range chunkPositions {
		vector, err := GetVector(location, realmSettings)
		if err != nil {
			return villages.NewVillage{}, err
		}
		current, err := chunks.NewChunk(ctx, db,
			villages.Coord{
				CoordX: vector.BaseX + (vector.DirectionX * position.CoordX * realmSettings.ChunkSize),
				CoordY: vector.BaseY + (vector.DirectionY * position.CoordY * realmSettings.ChunkSize),
			},
			realmSettings)
		if err != nil {
			return villages.NewVillage{}, err
		}
		if current.Population() < targetPopulation {
			coords, err := current.GetValidNewVillageCoords()
			if err != nil {
				return villages.NewVillage{}, nil
			}
			newVillage := villages.NewVillage{
				CoordX:   coords.CoordX,
				CoordY:   coords.CoordY,
				PlayerId: playerId,
			}
			err = villages.CreateVillage(ctx, db, newVillage)
			if err != nil {
				return villages.NewVillage{}, nil
			}
			return newVillage, nil
		}
	}
	return villages.NewVillage{}, nil
}
