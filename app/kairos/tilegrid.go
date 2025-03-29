package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jrecuero/ebiplay/pkg/engine"
)

type TileGrid struct {
	*engine.TileGrid
}

func NewTileGrid(tilemap *engine.TilemapJSON) *TileGrid {
	return &TileGrid{
		TileGrid: engine.NewTileGrid("main tilegrid", tilemap),
	}
}

func (t *TileGrid) ControlEntity(args ...any) {
	//x, y := entity.GetPos()
	//tileX, tileY := t.GetTilePosFromScreenPos(x, y)
	//key := GetKeyIntToString(tileX, tileY)
	//_ = key
	actor := args[0].(engine.IEntity)
	tilemapWidthInPixels, tilemapHeightInPixels := t.GetTilemapJSON().GetTilemapSizeInPixels()
	a := actor.(engine.IGridActor)
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		a.MoveUpdate("right", tilemapWidthInPixels, tilemapHeightInPixels)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		a.MoveUpdate("left", tilemapWidthInPixels, tilemapHeightInPixels)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		a.MoveUpdate("up", tilemapWidthInPixels, tilemapHeightInPixels)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		a.MoveUpdate("down", tilemapWidthInPixels, tilemapHeightInPixels)
	}
}

func (t *TileGrid) Update(args ...any) error {
	t.ControlEntity(args...)
	return t.TileGrid.Update(args...)
}
