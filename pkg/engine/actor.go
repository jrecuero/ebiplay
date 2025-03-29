package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type IActor interface {
	ISolidEntity
	Draw(*ebiten.Image, *Camera)
	GetBounds() image.Rectangle
	GetDx() float64
	GetDy() float64
	GetScale() float64
	GetSpeed() float64
	GetSpriteSheet() *SpriteSheet
	IsSolid() bool
	SetDx(float64) *Actor
	SetDy(float64) *Actor
	SetScale(float64) *Actor
	SetSpeed(float64) *Actor
	Update(...any) error
}

type Actor struct {
	*SolidEntity
	scale         float64
	speed, dx, dy float64
	spritesheet   *SpriteSheet
}

func IsInsideTilemapBoundary(x, y, width, height, tileWidth, tileHeight float64) bool {
	return (x >= 0) && (x < width-tileWidth) && (y >= 0) && (y < height-tileHeight)
}

func NewActor(name string, spritesheet *SpriteSheet, x, y float64) *Actor {
	return &Actor{
		SolidEntity: NewSolidEntity(name, x, y, 0, 0),
		spritesheet: spritesheet,
		scale:       1.0,
		speed:       2.0,
		dx:          0.0,
		dy:          0.0,
	}
}

func (a *Actor) ColorDraw(screen *ebiten.Image, camera *Camera) {
	image := a.GetSpriteSheet().GetFrameFor(a.GetSpriteSheet().frameType)
	ops := &colorm.DrawImageOptions{}
	ops.GeoM.Scale(a.GetScale(), a.GetScale())
	ops.GeoM.Translate(a.GetPos())
	if camera != nil {
		ops.GeoM.Translate(float64(camera.X), float64(camera.Y))
	}
	cm := colorm.ColorM{}
	cm.Translate(1.0, 1.0, 0.0, 0.5)
	colorm.DrawImage(screen, image, cm, ops)
}

func (a *Actor) Draw(screen *ebiten.Image, camera *Camera) {
	image := a.GetSpriteSheet().GetFrameFor(a.GetSpriteSheet().frameType)
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(a.GetScale(), a.GetScale())
	ops.GeoM.Translate(a.GetPos())
	if camera != nil {
		ops.GeoM.Translate(float64(camera.X), float64(camera.Y))
	}
	screen.DrawImage(image, ops)
}

func (a *Actor) GetBounds() image.Rectangle {
	if a.spritesheet != nil && a.spritesheet.Image != nil {
		maxX := int(a.x + float64(a.spritesheet.Width)*a.scale)
		maxY := int(a.y + float64(a.spritesheet.Height)*a.scale)
		rect := image.Rectangle{
			Min: image.Pt(int(a.x), int(a.y)),
			Max: image.Pt(maxX, maxY),
		}
		return rect
	}
	return image.Rectangle{}
}

func (a *Actor) GetDx() float64 {
	return a.dx
}

func (a *Actor) GetDy() float64 {
	return a.dy
}

func (a *Actor) GetScale() float64 {
	return a.scale
}

func (a *Actor) GetSpeed() float64 {
	return a.speed
}

func (a *Actor) GetSpriteSheet() *SpriteSheet {
	return a.spritesheet
}

func (a *Actor) IsSolid() bool {
	return true
}

func (a *Actor) SetScale(scale float64) *Actor {
	a.scale = scale
	return a
}

func (a *Actor) SetSpeed(speed float64) *Actor {
	a.speed = speed
	return a
}

func (a *Actor) SetDx(dx float64) *Actor {
	a.dx = dx
	return a
}

func (a *Actor) SetDy(dy float64) *Actor {
	a.dy = dy
	return a
}

func (a *Actor) Update(args ...any) error {
	tilemapWidthInPixels := args[0].(float64)
	tilemapHeightInPixels := args[1].(float64)
	x, y := a.GetPos()
	w := float64(a.GetSpriteSheet().Width) * a.GetScale()
	h := float64(a.GetSpriteSheet().Height) * a.GetScale()
	a.SetDx(0.0)
	a.SetDy(0.0)
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if IsInsideTilemapBoundary(x+a.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDx(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("right")
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if IsInsideTilemapBoundary(x-a.GetSpeed(), y, tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDx(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("left")
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		if IsInsideTilemapBoundary(x, y-a.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDy(-a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("up")
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if IsInsideTilemapBoundary(x, y+a.GetSpeed(), tilemapWidthInPixels, tilemapHeightInPixels, w, h) {
			a.SetDy(a.GetSpeed())
		}
		a.GetSpriteSheet().UpdateFrameType("down")
	}
	a.SetPos(x+a.GetDx(), y+a.GetDy())
	return nil
}

var _ IActor = (*Actor)(nil)
var _ IDrawable = (*Actor)(nil)
var _ IUpdatable = (*Actor)(nil)
