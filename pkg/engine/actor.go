package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IActor interface {
	Draw(*ebiten.Image, *Camera)
	GetName() string
	GetPos() (float64, float64)
	GetSpriteSheet() *SpriteSheet
	SetPos(float64, float64)
	Update(...any) error
}

type Actor struct {
	name        string
	x           float64
	y           float64
	spritesheet *SpriteSheet
}

func NewActor(name string, spritesheet *SpriteSheet, x, y float64) *Actor {
	return &Actor{
		name:        name,
		spritesheet: spritesheet,
		x:           x,
		y:           y,
	}
}

func (a *Actor) Draw(*ebiten.Image, *Camera) {
}

func (a *Actor) GetName() string {
	return a.name
}

func (a *Actor) GetPos() (float64, float64) {
	return a.x, a.y
}

func (a *Actor) GetSpriteSheet() *SpriteSheet {
	return a.spritesheet
}

func (a *Actor) SetPos(x, y float64) {
	a.x, a.y = x, y
}

func (a *Actor) Update(...any) error {
	return nil
}

var _ IActor = (*Actor)(nil)
