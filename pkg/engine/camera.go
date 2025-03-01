package engine

import "math"

type Camera struct {
	X, Y          float64
	Width, Height float64
}

func NewCamera(x, y, width, height float64) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (c *Camera) FollowTo(x, y float64) {
	c.X = -x + c.Width/2
	c.Y = -y + c.Height/2
}

func (c *Camera) Constrain(tilemapWidthInPixels, tilemapHeightInPixels float64) {
	c.X = math.Min(c.X, 0)
	c.Y = math.Min(c.Y, 0)

	c.X = math.Max(c.X, c.Width-tilemapWidthInPixels)
	c.Y = math.Max(c.Y, c.Height-tilemapHeightInPixels)
}
