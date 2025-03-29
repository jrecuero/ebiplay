package engine

import "image"

type ICollider interface {
	CollideWith(IEntity)
	GetBounds() image.Rectangle
	IsSolid() bool
}

func CheckCollidable(obj any) (ICollider, bool) {
	result, ok := obj.(ICollider)
	return result, ok
}
