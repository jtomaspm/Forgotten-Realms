package game

import (
	"fmt"
	"time"
)

type WorldMap struct {
	Map      map[Coordinates]*Village
	Villages []*Village
}

func NewWorldMap() *WorldMap {
	return &WorldMap{
		Map:      make(map[Coordinates]*Village),
		Villages: []*Village{},
	}
}

func (wm *WorldMap) AddVillage(village *Village) error {
	if _, exists := wm.Map[village.Coordinates]; exists {
		return fmt.Errorf("Village at %s already exists", village.Coordinates)
	}
	wm.Map[village.Coordinates] = village
	wm.Villages = append(wm.Villages, village)
	return nil
}

func (wm *WorldMap) FillWithVillages(max int) {
	i := 0
	for x := -500; x <= 500; x++ {
		for y := -500; y <= 500; y++ {
			if i >= max {
				return
			}
			coords := Coordinates{
				X: x,
				Y: y,
			}
			village := NewVillage(coords)
			wm.AddVillage(village)
			i++
		}
	}
}

func (wm *WorldMap) Update(deltaT time.Duration) {
	for _, village := range wm.Villages {
		village.Update(deltaT)
	}
}
