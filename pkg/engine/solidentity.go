package engine

import "image"

type ISolidEntity interface {
	IEntity
	ICollider
}

type SolidEntity struct {
	*Entity
}

func NewSolidEntity(name string, x, y float64, w, h int) *SolidEntity {
	return &SolidEntity{
		Entity: NewEntity(name, x, y, w, h),
	}
}

func (e *SolidEntity) CollideWith(IEntity) {
}

func (e *SolidEntity) GetBounds() image.Rectangle {
	maxX := int(e.x + float64(e.width))
	maxY := int(e.y + float64(e.height))
	return image.Rectangle{
		Min: image.Pt(int(e.x), int(e.y)),
		Max: image.Pt(maxX, maxY),
	}
}

func (e *SolidEntity) IsSolid() bool {
	return true
}

var _ IEntity = (*SolidEntity)(nil)
var _ ICollider = (*SolidEntity)(nil)
