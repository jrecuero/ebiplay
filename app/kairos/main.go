package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jrecuero/ebiplay/pkg/engine"
)

const (
	actorSpeed = 16

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
	Tilemap    *engine.TilemapJSON // all tile sprites.
	Actors     []engine.IActor     // all game actors.
	Camera     *engine.Camera
	colliders  []engine.ICollider
	tilegrid   *TileGrid
	keyhandler *engine.KeyboardHandler
}

func (g *Game) Update() error {
	g.keyhandler.Update()
	g.tilegrid.Update(g.Actors[0])
	for _, actor := range g.Actors {
		//actor.Update(tilemapWidthInPixels, tilemapHeightInPixels)
		if actor.GetName() == "knight" {
			g.Camera.FollowTo(actor.GetPos())
			g.Camera.Constrain(tilemapWidthInPixels, tilemapHeightInPixels)
		}
	}
	//g.tilegrid.ControlEntity(tilemapWidthInPixels, tilemapHeightInPixels, g.Actors[0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Tilemap.Draw(screen, g.Camera)
	//for _, actor := range g.Actors {
	//    actor.Draw(screen, g.Camera)
	//}
	g.tilegrid.Draw(screen, g.Camera)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	wd := "./"
	tilemapPath := filepath.Join(wd, "assets/tilemaps/tilemap.tmj")
	tilemapSpriteSheetPath := filepath.Join(wd, "assets/tilemaps/tileset.tsj")

	tilemap := engine.NewTilemapJSON(tilemapPath, tilemapSpriteSheetPath)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("TileMap Demo")

	knightPath := filepath.Join(wd, "assets/images/knight_walk.png")
	knightSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(knightPath)
	if err != nil {
		log.Fatalf("can not load knight images file %s: %s", knightPath, err)
	}
	knightSpriteSheet := engine.NewSpriteSheet(knightSpriteSheetImage, 4, 4, 16, 16)
	knightSpriteSheet.SetFrameMap(map[string][]int{
		"down":  {0, 0, 0, 16, 0, 32, 0, 48},
		"up":    {16, 0, 16, 16, 16, 32, 16, 48},
		"left":  {32, 0, 32, 16, 32, 32, 32, 48},
		"right": {48, 0, 48, 16, 48, 32, 48, 48},
	})
	knightSpriteSheet.SetFrameType("down")
	knightSpriteSheet.SetFrameSpeed(15)
	knight := NewKnight("knight", knightSpriteSheet, 0, 0)
	knight.SetSpeed(actorSpeed)

	spiritPath := filepath.Join(wd, "assets/images/spirit_walk.png")
	spiritSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(spiritPath)
	if err != nil {
		log.Fatalf("can not load spirit images file %s: %s", spiritPath, err)
	}
	spiritSpriteSheet := engine.NewSpriteSheet(spiritSpriteSheetImage, 4, 4, 16, 16)
	spiritSpriteSheet.SetFrameMap(map[string][]int{
		"down":  {0, 0, 0, 16, 0, 32, 0, 48},
		"up":    {16, 0, 16, 16, 16, 32, 16, 48},
		"left":  {32, 0, 32, 16, 32, 32, 32, 48},
		"right": {48, 0, 48, 16, 48, 32, 48, 48},
	})
	spiritSpriteSheet.SetFrameType("down")
	spiritSpriteSheet.SetFrameSpeed(15)
	spirit := NewSpirit("spirit", spiritSpriteSheet, 32, 32)

	g := &Game{
		Tilemap:    tilemap,
		Actors:     []engine.IActor{knight, spirit},
		Camera:     engine.NewCamera(0, 0, screenWidth, screenHeight),
		tilegrid:   NewTileGrid(tilemap),
		keyhandler: engine.NewKeyboardHandler("keyhandler"),
	}

	g.tilegrid.AddTileAt(0, 0, knight)
	g.tilegrid.AddTileAt(2, 2, spirit)

	for _, act := range g.Actors {
		g.colliders = append(g.colliders, act)
	}

	g.keyhandler.AddKeyBindingForKey(ebiten.KeyC, nil, func() {
		fmt.Println("ctrl-c was pressed")
		os.Exit(0)
	})

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
