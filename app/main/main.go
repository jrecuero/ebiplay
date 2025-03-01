package main

import (
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jrecuero/ebiplay/pkg/engine"
)

const (
	tilemapWidth  = 30
	tilemapHeight = 20
	tilemapSize   = tilemapWidth * tilemapHeight

	tileWidthInPixels  = 16
	tileHeightInPixels = 16

	tilemapWidthInPixels  float64 = tilemapWidth * tileWidthInPixels
	tilemapHeightInPixels float64 = tilemapHeight * tileHeightInPixels

	screenWidth  = (tilemapWidth * tileWidthInPixels) / 2
	screenHeight = (tilemapHeight * tileHeightInPixels) / 2
)

type Game struct {
	Tilemap *engine.TilemapJSON
	Actors  []engine.IActor
	Camera  *engine.Camera
}

func (g *Game) Update() error {
	for _, actor := range g.Actors {
		actor.Update(tilemapWidthInPixels, tilemapHeightInPixels)
		g.Camera.FollowTo(actor.GetPos())
		g.Camera.Constrain(tilemapWidthInPixels, tilemapHeightInPixels)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Tilemap.Draw(screen, g.Camera)
	for _, actor := range g.Actors {
		actor.Draw(screen, g.Camera)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//wd, err := os.Getwd()
	//if err != nil {
	//    log.Fatalf("error getting working directory: %s", err)
	//}
	wd := "app/main"
	tilemapPath := filepath.Join(wd, "assets/tilemaps/tilemap.tmj")
	imagePath := filepath.Join(wd, "assets/images/TilesetFloor.png")
	tilemap := engine.NewTilemapJSON(tilemapPath, imagePath)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("TileMap Demo")

	knightPath := filepath.Join(wd, "assets/images/Walk.png")
	knight := NewKnight("knight", knightPath, 16, 16)

	g := &Game{
		Tilemap: tilemap,
		Actors:  []engine.IActor{knight},
		Camera:  engine.NewCamera(0, 0, screenWidth, screenHeight),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
