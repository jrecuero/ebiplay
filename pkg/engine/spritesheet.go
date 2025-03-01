package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	Image   *ebiten.Image
	Rows    int
	Columns int
	Width   int
	Height  int
}

func NewSpriteSheet(image *ebiten.Image, rows, columns, width, height int) *SpriteSheet {
	return &SpriteSheet{
		Image:   image,
		Rows:    rows,
		Columns: columns,
		Width:   width,
		Height:  height,
	}
}

func (s *SpriteSheet) GetSpriteFromRowAndCol(row, col int) *ebiten.Image {
	x := col * s.Width
	y := row * s.Height
	return s.Image.SubImage(image.Rect(x, y, x+s.Width, y+s.Height)).(*ebiten.Image)
}
