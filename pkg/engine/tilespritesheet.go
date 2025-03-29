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
	path    string
	image   *ebiten.Image
	rows    int
	columns int
	width   int
	height  int
}

func NewTileSpriteSheet(jsonPath string) *TileSpriteSheet {
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("can not load tile sprite sheet json file %s: %s", jsonPath, err)
	}
	tilesetJSON := &TileSpriteSheetJSON{}
	if err := json.Unmarshal(content, tilesetJSON); err != nil {
		log.Fatalf("can not unmarshal tileset JSON file %s: %s", jsonPath, err)
	}

	// path to a JSON file that contains all information about the tile set.
	img, _, err := ebitenutil.NewImageFromFile(tilesetJSON.ImagePath)
	if err != nil {
		log.Fatalf("can not load tile sprite sheet image file %s: %s", jsonPath, err)
	}

	return &TileSpriteSheet{
		path:    jsonPath,
		image:   img,
		rows:    tilesetJSON.ImageHeight / tilesetJSON.TileHeight,
		columns: tilesetJSON.Columns,
		width:   tilesetJSON.TileWidth,
		height:  tilesetJSON.TileHeight,
	}
}

func (s *TileSpriteSheet) GetImage() *ebiten.Image {
	return s.image
}

func (s *TileSpriteSheet) GetSpriteForID(id int) *ebiten.Image {
	x := (id - 1) % s.columns
	y := (id - 1) / s.columns
	return s.image.SubImage(image.Rect(x*16, y*16, x*16+16, y*16+16)).(*ebiten.Image)
}

func (s *TileSpriteSheet) GetSpriteForRowAndCol(row, col int) *ebiten.Image {
	x := col * s.width
	y := row * s.height
	return s.image.SubImage(image.Rect(x, y, x+s.width, y+s.height)).(*ebiten.Image)
}

func (s *TileSpriteSheet) GetTileSize() (int, int) {
	return s.width, s.height
}
