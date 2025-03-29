package engine

import (
	"encoding/json"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	id_TILEMAP_MASK uint32 = 0x1FFFFFFF

	flipped_HORIZONTALLY_FLAG uint32 = 0x80000000
	flipped_VERTICALLY_FLAG   uint32 = 0x40000000
	flipped_DIAGONALLY_FLAG   uint32 = 0x20000000
	flipped_ALL_FLAG          uint32 = (flipped_HORIZONTALLY_FLAG | flipped_VERTICALLY_FLAG | flipped_DIAGONALLY_FLAG)
	op_TILEMAP_SHIFT                 = 28
)

// GetSpriteID function returns the sprite ID for a give tile in the map.
// Higher three bits contain some rotation operation done in the map with any
// given sprite.
func GetSpriteID(rawID uint32) uint32 {
	return rawID & id_TILEMAP_MASK
}

// GetSpriteOps function returns operation done in the map for a given sprite.
func GetSpriteOps(rawID uint32) uint32 {
	return (rawID & flipped_ALL_FLAG) >> op_TILEMAP_SHIFT
}

// DecodeTileID functions decodes the value in a tilemap returning the sprite
// ID and any operations done in the sprite in the tilemap. Operations are
// returned as a ebiten.GeoM instance.
func DecodeTileID(rawID uint32, tileWidth, tileHeight int, geoM *ebiten.GeoM) (uint32, *ebiten.GeoM) {
	id := GetSpriteID(rawID) // Remove flags

	// Diagonal flip (Swap X and Y)
	if rawID&flipped_DIAGONALLY_FLAG != 0 {
		geoM.Rotate(math.Pi / 2)              // Rotate 90° counterclockwise
		geoM.Translate(float64(tileWidth), 0) // Adjust position
	}

	// Flip Horizontal transformation
	if rawID&flipped_HORIZONTALLY_FLAG != 0 {
		geoM.Scale(-1, 1) // Flip horizontally
		geoM.Translate(float64(tileWidth), 0)
	}

	// Flip Vertical transformation
	if rawID&flipped_VERTICALLY_FLAG != 0 {
		geoM.Scale(1, -1) // Flip vertically
		geoM.Translate(0, float64(tileHeight))
	}

	return id, geoM
}

type TilemapLayerJSON struct {
	Data   []uint32 `json:"data"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

type TilemapJSON struct {
	Layers    []TilemapLayerJSON `json:"layers"`
	tileSheet *TileSpriteSheet
}

func NewTilemapJSON(tilemapPath string, tilesheetPath string) *TilemapJSON {
	content, err := os.ReadFile(tilemapPath)
	if err != nil {
		log.Fatalf("can not load tilemap JSON file %s: %s", tilemapPath, err)
	}

	tilesheet := NewTileSpriteSheet(tilesheetPath)
	if tilesheet == nil {
		log.Fatalf("can not load tile sheet sprite file %s: %s", tilesheetPath, err)
	}

	tilemapJSON := &TilemapJSON{}
	if err := json.Unmarshal(content, tilemapJSON); err != nil {
		log.Fatalf("can not unmarshal tilemap JSON file %s: %s", tilemapPath, err)
	}
	tilemapJSON.tileSheet = tilesheet

	return tilemapJSON
}

func (t *TilemapJSON) GetTilemapSize() (int, int) {
	if len(t.Layers) != 0 {
		return t.Layers[0].Width, t.Layers[0].Height
	}
	return 0, 0
}

func (t *TilemapJSON) GetTilemapSizeInPixels() (float64, float64) {
	tileWidth, tileHeight := t.GetTileSize()
	tilemapWidth, tilemapHeigt := t.GetTilemapSize()
	return float64(tilemapWidth * tileWidth), float64(tilemapHeigt * tileHeight)
}

func (t *TilemapJSON) GetTileSize() (int, int) {
	return t.tileSheet.GetTileSize()
}

func (t *TilemapJSON) Draw(screen *ebiten.Image, camera *Camera) {
	op := &ebiten.DrawImageOptions{}
	for _, layer := range t.Layers {
		for index, data := range layer.Data {
			w, h := t.tileSheet.GetTileSize()
			id, _ := DecodeTileID(data, w, h, &op.GeoM)
			tileImage := t.tileSheet.GetSpriteForID(int(id))
			screenX := (index % layer.Width) * w
			screenY := (index / layer.Width) * h
			op.GeoM.Translate(float64(screenX), float64(screenY))
			op.GeoM.Translate(float64(camera.X), float64(camera.Y))
			screen.DrawImage(tileImage, op)
			op.GeoM.Reset()
		}
	}
}
