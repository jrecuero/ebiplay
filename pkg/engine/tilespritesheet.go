package engine

import (
	"encoding/json"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileSpriteSheetJSON struct {
	Columns     int    `json:"columns"`
	ImagePath   string `json:"image"`
	ImageHeight int    `json:"imageheight"`
	ImageWidth  int    `json:"imagewidth"`
	Margin      int    `json:"margin"`
	TileCount   int    `json:"tilecount"`
	TileHeight  int    `json:"tileheight"`
	Tiles       []any  `json:"tiles"`
	TileWidth   int    `json:"tilewidth"`
}

type TileSpriteSheet struct {
	Path    string
	Image   *ebiten.Image
	Rows    int
	Columns int
	Width   int
	Height  int
}

func NewTileSpriteSheet(jsonPath string, rows, columns, width, height int) *TileSpriteSheet {
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("can not load tile sprite sheet json file %s: %s", jsonPath, err)
	}
	tilesetJSON := &TileSpriteSheetJSON{}
	if err := json.Unmarshal(content, tilesetJSON); err != nil {
		log.Fatalf("can not unmarshal tileset JSON file %s: %s", jsonPath, err)
	}

	// path to a JSON file that contains all information about the tile set.
	image, _, err := ebitenutil.NewImageFromFile(tilesetJSON.ImagePath)
	if err != nil {
		log.Fatalf("can not load tile sprite sheet image file %s: %s", jsonPath, err)
	}

	return &TileSpriteSheet{
		Path:    jsonPath,
		Image:   image,
		Rows:    tilesetJSON.ImageHeight / tilesetJSON.TileHeight,
		Columns: tilesetJSON.Columns,
		Width:   tilesetJSON.TileWidth,
		Height:  tilesetJSON.TileHeight,
	}
}

func (s *TileSpriteSheet) GetSpriteForID(id int) *ebiten.Image {
	x := (id - 1) % s.Columns
	y := (id - 1) / s.Columns
	return s.Image.SubImage(image.Rect(x*16, y*16, x*16+16, y*16+16)).(*ebiten.Image)
}

func (s *TileSpriteSheet) GetSpriteForRowAndCol(row, col int) *ebiten.Image {
	x := col * s.Width
	y := row * s.Height
	return s.Image.SubImage(image.Rect(x, y, x+s.Width, y+s.Height)).(*ebiten.Image)
}
