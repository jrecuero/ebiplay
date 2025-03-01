package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type IActor interface {
	Draw(*ebiten.Image, *Camera)
	GetName() string
	GetPos() (float64, float64)
	GetSpriteSheet() *ebiten.Image
	SetPos(float64, float64)
	Update(...any) error
}

type Actor struct {
	name        string
	x           float64
	y           float64
	spritesheet *ebiten.Image
}

func NewActor(name, imagesPath string, x, y float64) *Actor {
	spritesheet, _, err := ebitenutil.NewImageFromFile(imagesPath)
	if err != nil {
		log.Fatalf("can not load actor images file %s: %s", imagesPath, err)
	}
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

func (a *Actor) GetSpriteSheet() *ebiten.Image {
	return a.spritesheet
}

func (a *Actor) SetPos(x, y float64) {
	a.x, a.y = x, y
}

func (a *Actor) Update(...any) error {
	return nil
}

var _ IActor = (*Actor)(nil)
