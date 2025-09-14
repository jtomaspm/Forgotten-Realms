package simulation

import (
	"log"
	"time"
)

type Tick struct {
	tick   *time.Ticker
	now    time.Time
	deltaT time.Duration
}

func NewTick(framerate int) *Tick {
	deltaT := time.Second / time.Duration(framerate)
	return &Tick{
		tick:   time.NewTicker(deltaT),
		now:    time.Now(),
		deltaT: deltaT,
	}
}

func (t *Tick) Start() {
	t.tick.Reset(t.deltaT)
	t.now = time.Now()
}

func (t Tick) Now() time.Time {
	return t.now
}

func (t Tick) DeltaT() time.Duration {
	return t.deltaT
}

func (t *Tick) NextFrame() {
	log.Printf("Frame Time: %d ms\n", time.Since(t.now).Milliseconds())
	<-t.tick.C
	t.now = time.Now()
}
