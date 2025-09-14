package simulation

import "world_server/pkg/game"

type Simulation struct {
	tick      *Tick
	world_map *game.WorldMap
}

func New() Simulation {
	return Simulation{
		tick:      NewTick(FRAMERATE),
		world_map: game.NewWorldMap(),
	}
}

func (s *Simulation) Start() {
	s.tick.Start()
	s.world_map.FillWithVillages(10)
	for {
		s.world_map.Update(s.tick.DeltaT())
		s.tick.NextFrame()
	}
}
