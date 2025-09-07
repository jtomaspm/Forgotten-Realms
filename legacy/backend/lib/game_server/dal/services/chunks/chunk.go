package chunks

import (
	"backend/lib/game_server/dal/services/settings"
	"backend/lib/game_server/dal/services/villages"
	"backend/pkg/database"
	"context"
	"fmt"
	"log"
	"math/rand"
)

type Chunk struct {
	CoordX   int                `json:"coord_x"`
	CoordY   int                `json:"coord_y"`
	Size     int                `json:"size"`
	Villages []villages.Village `json:"villages"`
}

func NewChunk(ctx context.Context, db database.Querier, coords villages.Coords, realmSettings settings.Realm) (Chunk, error) {
	if !(coords.CoordX < realmSettings.MapSize || coords.CoordX >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordX")
	}
	if !(coords.CoordY < realmSettings.MapSize || coords.CoordY >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordY")
	}
	var chunk Chunk
	log.Println("chunkCoords: ", coords)
	chunk.CoordX = coords.CoordX - (coords.CoordX % realmSettings.ChunkSize)
	chunk.CoordY = coords.CoordY - (coords.CoordY % realmSettings.ChunkSize)
	chunk.Size = realmSettings.ChunkSize
	v, err := villages.GetVillagesInRange(
		ctx,
		db,
		chunk.CoordX,
		chunk.CoordY,
		chunk.CoordX+realmSettings.ChunkSize,
		chunk.CoordY+realmSettings.ChunkSize,
		realmSettings)
	if err != nil {
		return Chunk{}, err
	}
	chunk.Villages = v
	return chunk, nil
}

func (c *Chunk) Population() float32 {
	return float32(len(c.Villages) / c.Size)
}

func (c *Chunk) GetValidNewVillageCoords() (villages.Coords, error) {
	if c.Population() >= 1 {
		return villages.Coords{}, fmt.Errorf("maximum population reached")
	}
	usedCoords := make(map[string]bool)
	for _, village := range c.Villages {
		key := fmt.Sprintf("%d,%d", village.CoordX, village.CoordY)
		usedCoords[key] = true
	}

	maxAttempts := c.Size * c.Size
	for range maxAttempts {
		x := rand.Intn(c.Size) + c.CoordX
		y := rand.Intn(c.Size) + c.CoordY
		key := fmt.Sprintf("%d,%d", x, y)
		if !usedCoords[key] {
			return villages.Coords{CoordX: x, CoordY: y}, nil
		}
	}
	return villages.Coords{}, fmt.Errorf("could not find valid coordinates after %d attempts", maxAttempts)
}

func (c *Chunk) GetTop(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coords{
		CoordX: c.CoordX,
		CoordY: c.CoordY - rs.ChunkSize,
	}, rs)
}

func (c *Chunk) GetBottom(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coords{
		CoordX: c.CoordX,
		CoordY: c.CoordY + rs.ChunkSize,
	}, rs)
}

func (c *Chunk) GetLeft(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coords{
		CoordX: c.CoordX - rs.ChunkSize,
		CoordY: c.CoordY,
	}, rs)
}

func (c *Chunk) GetRight(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coords{
		CoordX: c.CoordX + rs.ChunkSize,
		CoordY: c.CoordY,
	}, rs)
}
