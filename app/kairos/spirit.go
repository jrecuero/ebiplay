package main

import (
	"github.com/jrecuero/ebiplay/pkg/engine"
)

type Spirit struct {
	*engine.Actor
}

func NewSpirit(name string, spritesheet *engine.SpriteSheet, x, y float64) *Spirit {
	spirit := &Spirit{
		Actor: engine.NewActor(name, spritesheet, x, y),
	}
	return spirit
}

func (s *Spirit) Update(args ...any) error {
	return nil
}

var _ engine.IActor = (*Spirit)(nil)
