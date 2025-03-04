package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	Tilemap            *engine.TilemapJSON
	Actors             []engine.IActor
	Camera             *engine.Camera
	warriorSpriteSheet *engine.SpriteSheet
	counter            int
	frame              int
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
	//ops := &ebiten.DrawImageOptions{}
	//ops.GeoM.Scale(0.2, 0.2)
	//g.counter = (g.counter + 1) % 60
	//g.frame += 1
	//i := 1
	//if g.frame == 180 {
	//    g.frame = 0
	//}
	//if g.frame > 120 {
	//    i = 2
	//}
	//var img *ebiten.Image
	//index := 512 * i
	//if g.counter > 45 {
	//    img = g.warriorSpriteSheet.Image.SubImage(image.Rect(1536, index, 2048, index+512)).(*ebiten.Image)
	//} else if g.counter > 30 {
	//    img = g.warriorSpriteSheet.Image.SubImage(image.Rect(1024, index, 1536, index+512)).(*ebiten.Image)
	//} else if g.counter > 15 {
	//    img = g.warriorSpriteSheet.Image.SubImage(image.Rect(512, index, 1024, index+512)).(*ebiten.Image)
	//} else {
	//    img = g.warriorSpriteSheet.Image.SubImage(image.Rect(0, index, 512, index+512)).(*ebiten.Image)
	//}
	//screen.DrawImage(img, ops)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//wd, err := os.Getwd()
	//if err != nil {
	//    log.Fatalf("error getting working directory: %s", err)
	//}
	//wd := "app/main"
	wd := "./"
	tilemapPath := filepath.Join(wd, "assets/tilemaps/tilemap.tmj")
	tilemapSpriteSheetPath := filepath.Join(wd, "assets/tilemaps/tileset.tsj")
	imagePath := filepath.Join(wd, "assets/images/TilesetFloor.png")
	tilemap := engine.NewTilemapJSON(tilemapPath, imagePath)

	//ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowSize(screenWidth*4, screenHeight*4)
	ebiten.SetWindowTitle("TileMap Demo")

	knightPath := filepath.Join(wd, "assets/images/Walk.png")
	knightSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(knightPath)
	if err != nil {
		log.Fatalf("can not load knight images file %s: %s", knightPath, err)
	}
	knightSpriteSheet := engine.NewSpriteSheet(knightSpriteSheetImage, 4, 4, 16, 16)
	// add animation to the spritesheet
	knightSpriteSheet.FrameTypes = []string{"down", "up", "left", "right"}
	knightSpriteSheet.FrameType = "down"
	knightSpriteSheet.FrameMap = map[string][]int{
		"down":  {0, 0, 0, 16, 0, 32, 0, 48},
		"up":    {16, 0, 16, 16, 16, 32, 16, 48},
		"left":  {32, 0, 32, 16, 32, 32, 32, 48},
		"right": {48, 0, 48, 16, 48, 32, 48, 48},
	}
	knightSpriteSheet.FrameSpeed = 15
	knight := NewKnight("knight", knightSpriteSheet, 16, 16)
	_ = knight

	// male warrior
	warriorPath := filepath.Join(wd, "assets/images/male_warrior.png")
	warriorSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(warriorPath)
	if err != nil {
		log.Fatalf("can not load male warrior images file %s: %s", warriorPath, err)
	}
	warriorSpriteSheet := engine.NewSpriteSheet(warriorSpriteSheetImage, 16, 4, 512, 512)
	// add animation to the spritesheet
	warriorSpriteSheet.FrameTypes = []string{"down", "up", "left", "right", "attack/down", "attack/up", "attack/left", "attack/right"}
	warriorSpriteSheet.FrameType = "down"
	warriorSpriteSheet.FrameMap = map[string][]int{
		"down":         {0, 512, 512, 512, 1024, 512, 1536, 512},
		"up":           {0, 2560, 512, 2560, 1024, 2560, 1536, 2560},
		"left":         {0, 4608, 512, 4608, 1024, 4608, 1536, 4608},
		"right":        {0, 6656, 512, 6656, 1024, 6656, 1536, 6656},
		"attack/down":  {0, 1024, 512, 1024, 1024, 1024, 1536, 1024},
		"attack/up":    {0, 3072, 512, 3072, 1024, 3072, 1536, 3072},
		"attack/left":  {0, 5120, 512, 5120, 1024, 5120, 1536, 5120},
		"attack/right": {0, 7168, 512, 7168, 1024, 7168, 1536, 7168},
	}
	warriorSpriteSheet.FrameSpeed = 15
	warrior := NewWarrior("warrior", warriorSpriteSheet, 0, 0)
	_ = warrior

	tileSpriteSheet := engine.NewTileSpriteSheet(tilemapSpriteSheetPath, 0, 0, 16, 16)
	fmt.Printf("%#+v\n", tileSpriteSheet)

	g := &Game{
		Tilemap: tilemap,
		//Actors:  []engine.IActor{knight},
		Actors: []engine.IActor{warrior},
		Camera: engine.NewCamera(0, 0, screenWidth, screenHeight),
		//warriorSpriteSheet: warriorSpriteSheet,
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
