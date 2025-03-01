package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jrecuero/ebiplay/pkg/engine"
)

const (
	down int = iota
	up
	left
	right
)

type Knight struct {
	*engine.Actor
	dir      int
	indexes  map[int][]int
	tcounter int
	tframe   int
}

func NewKnight(name, imagesPath string, x, y float64) *Knight {
	return &Knight{
		Actor: engine.NewActor(name, imagesPath, x, y),
		dir:   down,
		indexes: map[int][]int{
			down:  {0, 0, 0, 16, 0, 32, 0, 48},
			up:    {16, 0, 16, 16, 16, 32, 16, 48},
			left:  {32, 0, 32, 16, 32, 32, 32, 48},
			right: {48, 0, 48, 16, 48, 32, 48, 48},
		},
		tcounter: 0,
		tframe:   0,
	}
}

func isInsideTilemapBoundary(x, y, width, height float64) bool {
	return (x >= 0) && (x < width-16) && (y >= 0) && (y < height-16)
}

func (k *Knight) Update(args ...any) error {
	tilemapWidthInPixels := args[0].(float64)
	tilemapHeightInPixels := args[1].(float64)
	speed := 2.0
	x, y := k.GetPos()
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if isInsideTilemapBoundary(x+speed, y, tilemapWidthInPixels, tilemapHeightInPixels) {
			k.SetPos(x+speed, y)
		}
		k.dir = right
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if isInsideTilemapBoundary(x-speed, y, tilemapWidthInPixels, tilemapHeightInPixels) {
			k.SetPos(x-speed, y)
		}
		k.dir = left
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if isInsideTilemapBoundary(x, y-speed, tilemapWidthInPixels, tilemapHeightInPixels) {
			k.SetPos(x, y-speed)
		}
		k.dir = up
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if isInsideTilemapBoundary(x, y+speed, tilemapWidthInPixels, tilemapHeightInPixels) {
			k.SetPos(x, y+speed)
		}
		k.dir = down
	}
	return nil
}

func (k *Knight) Draw(screen *ebiten.Image, camera *engine.Camera) {
	indexes := k.indexes[k.dir]
	if k.tcounter = (k.tcounter + 1) % 15; k.tcounter == 0 {
		k.tframe = (k.tframe + 1) % 4
	}
	i := k.tframe * 2
	x, y := indexes[i], indexes[i+1]
	image := k.GetSpriteSheet().SubImage(image.Rect(x, y, x+16, y+16)).(*ebiten.Image)
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(k.GetPos())
	ops.GeoM.Translate(float64(camera.X), float64(camera.Y))
	screen.DrawImage(image, ops)
}
