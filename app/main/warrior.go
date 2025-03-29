package main

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jrecuero/ebiplay/pkg/engine"
)

type Warrior struct {
	*engine.Actor
}

func NewWarrior(name string, spritesheet *engine.SpriteSheet, x, y float64) *Warrior {
	warrior := &Warrior{
		Actor: engine.NewActor(name, spritesheet, x, y),
	}
	warrior.SetScale(0.08).SetSpeed(2.0)
	return warrior
}

//func isWarriorInsideTilemapBoundary(x, y, width, height, tileWidth, tileHeight float64) bool {
//    return (x >= 0) && (x < width-tileWidth) && (y >= 0) && (y < height-tileHeight)
//}

func (k *Warrior) Update(args ...any) error {
	//tilemapWidthInPixels := args[0].(float64)
	//tilemapHeightInPixels := args[1].(float64)
	//x, y := k.GetPos()
	//w := float64(k.GetSpriteSheet().Width) * k.GetScale()
	//h := float64(k.GetSpriteSheet().Height) * k.GetScale()
	//if ebiten.IsKeyPressed(ebiten.KeyRight) {
	//    if engine.IsInsideTilemapBoundary(x+k.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
	//        k.SetPos(x+k.GetSpeed(), y)
	//    }
	//    k.GetSpriteSheet().UpdateFrameType("right")
	//} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	//    if engine.IsInsideTilemapBoundary(x-k.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
	//        k.SetPos(x-k.GetSpeed(), y)
	//    }
	//    k.GetSpriteSheet().UpdateFrameType("left")
	//} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
	//    if engine.IsInsideTilemapBoundary(x, y-k.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
	//        k.SetPos(x, y-k.GetSpeed())
	//    }
	//    k.GetSpriteSheet().UpdateFrameType("up")
	//} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
	//    if engine.IsInsideTilemapBoundary(x, y+k.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
	//        k.SetPos(x, y+k.GetSpeed())
	//    }
	//    k.GetSpriteSheet().UpdateFrameType("down")
	//} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
	//    frameType := k.GetSpriteSheet().FrameType
	//    if !strings.Contains(frameType, "attack") {
	//        newFrameType := fmt.Sprintf("attack/%s", frameType)
	//        k.GetSpriteSheet().UpdateFrameType(newFrameType)
	//    }
	//}
	if err := k.Actor.Update(args...); err != nil {
		return err
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		frameType := k.GetSpriteSheet().GetFrameType()
		if !strings.Contains(frameType, "attack") {
			newFrameType := fmt.Sprintf("attack/%s", frameType)
			k.GetSpriteSheet().UpdateFrameType(newFrameType)
		}
	}
	return nil
}

//func (k *Warrior) Draw(screen *ebiten.Image, camera *engine.Camera) {
//    image := k.GetSpriteSheet().GetFrameFor(k.GetSpriteSheet().FrameType)
//    ops := &ebiten.DrawImageOptions{}
//    ops.GeoM.Scale(k.GetScale(), k.GetScale())
//    ops.GeoM.Translate(k.GetPos())
//    ops.GeoM.Translate(float64(camera.X), float64(camera.Y))
//    screen.DrawImage(image, ops)
//}

var _ engine.IActor = (*Warrior)(nil)
