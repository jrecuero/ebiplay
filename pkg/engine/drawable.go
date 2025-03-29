package engine

import "github.com/hajimehoshi/ebiten/v2"

type IDrawable interface {
	Draw(*ebiten.Image, *Camera)
}

func CheckDrawable(obj any) (IDrawable, bool) {
	result, ok := obj.(IDrawable)
	return result, ok
}
