package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileSequence []IEntity

type TileGrid struct {
	*Base
	height  int
	tilemap *TilemapJSON
	tiles   map[string]TileSequence
	width   int
}

func NewTileGrid(name string, tilemap *TilemapJSON) *TileGrid {
	width, height := tilemap.GetTileSize()
	tilegrid := &TileGrid{
		Base:    NewBase(name),
		height:  height,
		tilemap: tilemap,
		tiles:   make(map[string]TileSequence),
		width:   width,
	}
	return tilegrid
}

func GetKeyIntToString(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func GetKeyStringToInt(key string) (int, int) {
	split := strings.Split(key, ":")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return x, y
}

func (t *TileGrid) AddTile(key string, tile IEntity) error {
	t.tiles[key] = append(t.tiles[key], tile)
	return nil
}

func (t *TileGrid) AddTileAt(x, y int, tile IEntity) error {
	key := GetKeyIntToString(x, y)
	return t.AddTile(key, tile)
}

func (t *TileGrid) Draw(screen *ebiten.Image, camera *Camera) {
	for _, tiles := range t.tiles {
		for _, tile := range tiles {
			if result, ok := CheckDrawable(tile); ok {
				result.Draw(screen, camera)
			}
		}
	}
}

func (t *TileGrid) GetScreenPosFromTilePos(tileX, tileY int) (float64, float64) {
	var x float64 = float64(tileX * t.width)
	var y float64 = float64(tileY * t.height)
	return x, y
}

func (t *TileGrid) GetTilemapJSON() *TilemapJSON {
	return t.tilemap
}

func (t *TileGrid) GetTiles() map[string]TileSequence {
	return t.tiles
}

func (t *TileGrid) GetTilesAt(key string) TileSequence {
	return t.tiles[key]
}

func (t *TileGrid) GetTilePosFromScreenPos(x, y float64) (int, int) {
	var tileX int = int(x) / t.width
	var tileY int = int(y) / t.height
	return tileX, tileY
}

func (t *TileGrid) IsTileAt(key string, tile IEntity) bool {
	if tiles := t.GetTilesAt(key); tiles != nil {
		for _, atile := range tiles {
			if atile == tile {
				return true
			}
		}
	}
	return false
}

func (t *TileGrid) MoveTileFromTo(from, to string, tile IEntity) error {
	if err := t.RemoveTile(from, tile); err != nil {
		return err
	}
	return t.AddTile(to, tile)
}

func (t *TileGrid) RemoveTile(key string, tile IEntity) error {
	if tiles, ok := t.tiles[key]; ok {
		for i, atile := range tiles {
			if atile == tile {
				t.tiles[key] = append(t.tiles[key][:i], t.tiles[key][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (t *TileGrid) Update(args ...any) error {
	for _, tiles := range t.tiles {
		for _, tile := range tiles {
			if result, ok := CheckUpdatable(tile); ok {
				return result.Update(args...)
			}
		}
	}
	return nil
}
