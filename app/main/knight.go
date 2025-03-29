package main

import (
	"github.com/jrecuero/ebiplay/pkg/engine"
)

type Knight struct {
	*engine.Actor
}

func NewKnight(name string, spritesheet *engine.SpriteSheet, x, y float64) *Knight {
	return &Knight{
		Actor: engine.NewActor(name, spritesheet, x, y),
	}
}

//func isInsideTilemapBoundary(x, y, width, height, tileWidth, tileHeight float64) bool {
//    return (x >= 0) && (x < width-tileWidth) && (y >= 0) && (y < height-tileHeight)
//}

//func (k *Knight) Update(args ...any) error {
//    tilemapWidthInPixels := args[0].(float64)
//    tilemapHeightInPixels := args[1].(float64)
//    speed := 2.0
//    x, y := k.GetPos()
//    w := float64(k.GetSpriteSheet().Width)
//    h := float64(k.GetSpriteSheet().Height)
//    if ebiten.IsKeyPressed(ebiten.KeyRight) {
//        if engine.IsInsideTilemapBoundary(x+speed, y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
//            k.SetPos(x+speed, y)
//        }
//        k.GetSpriteSheet().UpdateFrameType("right")
//    } else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
//        if engine.IsInsideTilemapBoundary(x-speed, y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
//            k.SetPos(x-speed, y)
//        }
//        k.GetSpriteSheet().UpdateFrameType("left")
//    } else if ebiten.IsKeyPressed(ebiten.KeyUp) {
//        if engine.IsInsideTilemapBoundary(x, y-speed, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
//            k.SetPos(x, y-speed)
//        }
//        k.GetSpriteSheet().UpdateFrameType("up")
//    } else if ebiten.IsKeyPressed(ebiten.KeyDown) {
//        if engine.IsInsideTilemapBoundary(x, y+speed, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
//            k.SetPos(x, y+speed)
//        }
//        k.GetSpriteSheet().UpdateFrameType("down")
//    }
//    return nil
//}

//func (k *Knight) Draw(screen *ebiten.Image, camera *engine.Camera) {
//    image := k.GetSpriteSheet().GetFrameFor(k.GetSpriteSheet().FrameType)
//    ops := &ebiten.DrawImageOptions{}
//    ops.GeoM.Translate(k.GetPos())
//    ops.GeoM.Translate(float64(camera.X), float64(camera.Y))
//    screen.DrawImage(image, ops)
//}

var _ engine.IActor = (*Knight)(nil)
