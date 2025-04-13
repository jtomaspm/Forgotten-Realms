package chunks

import (
	"backend/lib/game_server/dal/services/settings"
	"backend/lib/game_server/dal/services/villages"
	"backend/pkg/database"
	"context"
	"fmt"
)

type Chunk struct {
	CoordX   int                `json:"coord_x"`
	CoordY   int                `json:"coord_y"`
	Size     int                `json:"size"`
	Villages []villages.Village `json:"villages"`
}

func NewChunk(ctx context.Context, db database.Querier, coord villages.Coord, realmSettings settings.Realm) (Chunk, error) {
	if !(coord.CoordX < realmSettings.MapSize || coord.CoordX >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordX")
	}
	if !(coord.CoordY < realmSettings.MapSize || coord.CoordY >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordY")
	}
	var chunk Chunk
	chunk.CoordX = chunk.CoordX - (chunk.CoordX % realmSettings.ChunkSize)
	chunk.CoordY = chunk.CoordY - (chunk.CoordY % realmSettings.ChunkSize)
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

func (c *Chunk) GetValidNewVillageCoords() (villages.Coord, error) {
	if c.Population() >= 1 {
		return villages.Coord{}, fmt.Errorf("maximum population reached")
	}
	x := 1
	y := 1
	return villages.Coord{
		CoordX: x,
		CoordY: y,
	}, nil
}

func (c *Chunk) GetTop(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coord{
		CoordX: c.CoordX,
		CoordY: c.CoordY - rs.ChunkSize,
	}, rs)
}

func (c *Chunk) GetBottom(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coord{
		CoordX: c.CoordX,
		CoordY: c.CoordY + rs.ChunkSize,
	}, rs)
}

func (c *Chunk) GetLeft(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coord{
		CoordX: c.CoordX - rs.ChunkSize,
		CoordY: c.CoordY,
	}, rs)
}

func (c *Chunk) GetRight(ctx context.Context, db database.Querier, rs settings.Realm) (Chunk, error) {
	return NewChunk(ctx, db, villages.Coord{
		CoordX: c.CoordX + rs.ChunkSize,
		CoordY: c.CoordY,
	}, rs)
}
