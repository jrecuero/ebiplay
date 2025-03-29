package main

import (
	"github.com/jrecuero/ebiplay/pkg/engine"
)

type Knight struct {
	*engine.GridActor
}

func NewKnight(name string, spritesheet *engine.SpriteSheet, x, y float64) *Knight {
	return &Knight{
		GridActor: engine.NewGridActor(name, spritesheet, x, y),
	}
}
func (a *Knight) Update(args ...any) error {
	return nil
}

var _ engine.IActor = (*Knight)(nil)
