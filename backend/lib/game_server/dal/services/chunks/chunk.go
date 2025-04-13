package chunks

import (
	"backend/lib/game_server/configuration"
	"backend/lib/game_server/dal/services/villages"
	"backend/pkg/database"
	"context"
	"fmt"
)

type Coord struct {
	CoordX int `json:"coord_x"`
	CoordY int `json:"coord_y"`
}

type Chunk struct {
	CoordX   int                `json:"coord_x"`
	CoordY   int                `json:"coord_y"`
	Size     int                `json:"size"`
	Villages []villages.Village `json:"villages"`
}

func NewChunk(ctx context.Context, db database.Querier, coord Coord) (Chunk, error) {
	if !(coord.CoordX < configuration.MAP_SIZE || coord.CoordX >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordX")
	}
	if !(coord.CoordY < configuration.MAP_SIZE || coord.CoordY >= 0) {
		return Chunk{}, fmt.Errorf("invalid CoordY")
	}
	var chunk Chunk
	chunk.CoordX = chunk.CoordX - (chunk.CoordX % configuration.CHUNK_SIZE)
	chunk.CoordY = chunk.CoordY - (chunk.CoordY % configuration.CHUNK_SIZE)
	chunk.Size = configuration.CHUNK_SIZE
	v, err := villages.GetVillagesInRange(ctx, db, chunk.CoordX, chunk.CoordY, chunk.CoordX+configuration.CHUNK_SIZE, chunk.CoordY+configuration.CHUNK_SIZE)
	if err != nil {
		return Chunk{}, err
	}
	chunk.Villages = v
	return chunk, nil
}

func (c *Chunk) GetTop(ctx context.Context, db database.Querier) (Chunk, error) {
	return NewChunk(ctx, db, Coord{
		CoordX: c.CoordX,
		CoordY: c.CoordY - configuration.CHUNK_SIZE,
	})
}

func (c *Chunk) GetBottom(ctx context.Context, db database.Querier) (Chunk, error) {
	return NewChunk(ctx, db, Coord{
		CoordX: c.CoordX,
		CoordY: c.CoordY + configuration.CHUNK_SIZE,
	})
}

func (c *Chunk) GetLeft(ctx context.Context, db database.Querier) (Chunk, error) {
	return NewChunk(ctx, db, Coord{
		CoordX: c.CoordX - configuration.CHUNK_SIZE,
		CoordY: c.CoordY,
	})
}

func (c *Chunk) GetRight(ctx context.Context, db database.Querier) (Chunk, error) {
	return NewChunk(ctx, db, Coord{
		CoordX: c.CoordX + configuration.CHUNK_SIZE,
		CoordY: c.CoordY,
	})
}
