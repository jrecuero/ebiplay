package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IGridActor interface {
	MoveUpdate(string, float64, float64) error
}

type GridActor struct {
	*Actor
}

func NewGridActor(name string, spritesheet *SpriteSheet, x, y float64) *GridActor {
	return &GridActor{
		Actor: NewActor(name, spritesheet, x, y),
	}
}

func (a *GridActor) MoveUpdate(moveto string, width, height float64) error {
	x, y := a.GetPos()
	w := float64(a.GetSpriteSheet().Width) * a.GetScale()
	h := float64(a.GetSpriteSheet().Height) * a.GetScale()
	a.SetDx(0.0)
	a.SetDy(0.0)
	switch moveto {
	case "right":
		if IsInsideTilemapBoundary(x+a.GetSpeed(), y, width, height, w, h) {
			a.SetDx(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("right")
	case "left":
		if IsInsideTilemapBoundary(x-a.GetSpeed(), y, width, height, w, h) {
			a.SetDx(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("left")
	case "up":
		if IsInsideTilemapBoundary(x, y-a.GetSpeed(), width, height, w, h) {
			a.SetDy(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("up")
	case "down":
		if IsInsideTilemapBoundary(x, y+a.GetSpeed(), width, height, w, h) {
			a.SetDy(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("down")
	default:
		return fmt.Errorf("unknown movement direction %s", moveto)
	}
	a.SetPos(x+a.GetDx(), y+a.GetDy())
	return nil
}

func (a *GridActor) Update(args ...any) error {
	tilemapWidthInPixels := args[0].(float64)
	tilemapHeightInPixels := args[1].(float64)
	x, y := a.GetPos()
	w := float64(a.GetSpriteSheet().Width) * a.GetScale()
	h := float64(a.GetSpriteSheet().Height) * a.GetScale()
	a.SetDx(0.0)
	a.SetDy(0.0)
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if IsInsideTilemapBoundary(x+a.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDx(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("right")
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if IsInsideTilemapBoundary(x-a.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDx(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("left")
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if IsInsideTilemapBoundary(x, y-a.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDy(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("up")
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if IsInsideTilemapBoundary(x, y+a.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDy(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("down")
	}
	a.SetPos(x+a.GetDx(), y+a.GetDy())
	return nil
}

var _ IGridActor = (*GridActor)(nil)
