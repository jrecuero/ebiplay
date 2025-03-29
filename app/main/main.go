package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

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

type Event struct {
	*engine.SolidEntity
	*engine.Event
}

func NewEvent(name string, x, y float64, w, h int) *Event {
	return &Event{
		SolidEntity: engine.NewSolidEntity(name, x, y, w, h),
		Event:       &engine.Event{},
	}
}

func (e *Event) CollideWith(other engine.IEntity) {
	e.HandleEvent(other)
}

func (e *Event) HandleEvent(args ...any) error {
	if !e.IsTriggered() {
		e.SetTriggered(true)
		other := args[0].(engine.IEntity)
		fmt.Printf("Handle event %s collide with %s\n", e.GetName(), other.GetName())
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Printf("reset event %s\n", e.GetName())
			e.SetTriggered(false)
		}()
	}
	return nil
}

type Game struct {
	Tilemap   *engine.TilemapJSON
	Actors    []engine.IActor
	Camera    *engine.Camera
	menu      *engine.Menu
	colliders []engine.ICollider
}

func checkHorizontalCollision(actor engine.IActor, colliders []engine.ICollider) {
	x, y := actor.GetPos()
	dx := actor.GetDx()
	for _, coll := range colliders {
		if collActor, ok := coll.(engine.IActor); ok && actor == collActor {
			continue
		}
		if actor.GetBounds().Overlaps(coll.GetBounds()) {
			actor.CollideWith(coll.(engine.IEntity))
			coll.CollideWith(actor)
			actor.SetPos(x-dx, y)
		}
	}
}

func checkVerticalCollision(actor engine.IActor, colliders []engine.ICollider) {
	x, y := actor.GetPos()
	dy := actor.GetDy()
	for _, coll := range colliders {
		if collActor, ok := coll.(engine.IActor); ok && actor == collActor {
			continue
		}
		if actor.GetBounds().Overlaps(coll.GetBounds()) {
			fmt.Printf("collision %s with %s\n", actor.GetName(), coll.(engine.IEntity).GetName())
			actor.CollideWith(coll.(engine.IEntity))
			coll.CollideWith(actor)
			actor.SetPos(x, y-dy)
		}
	}
}

func (g *Game) Update() error {
	for _, actor := range g.Actors {
		actor.Update(tilemapWidthInPixels, tilemapHeightInPixels)
		if actor.GetName() == "knight" {
			g.Camera.FollowTo(actor.GetPos())
			g.Camera.Constrain(tilemapWidthInPixels, tilemapHeightInPixels)
		}

		checkHorizontalCollision(actor, g.colliders)

		checkVerticalCollision(actor, g.colliders)
	}
	//g.menu.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Tilemap.Draw(screen, g.Camera)
	for _, actor := range g.Actors {
		actor.Draw(screen, g.Camera)
	}
	g.menu.Draw(screen)
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
	//imagePath := filepath.Join(wd, "assets/images/TilesetFloor.png")

	//tileSpriteSheet := engine.NewTileSpriteSheet(tilemapSpriteSheetPath)
	//fmt.Printf("%#+v\n", tileSpriteSheet)

	tilemap := engine.NewTilemapJSON(tilemapPath, tilemapSpriteSheetPath)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("TileMap Demo")

	knightPath := filepath.Join(wd, "assets/images/Walk.png")
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

	spiritPath := filepath.Join(wd, "assets/images/spirit_idle.png")
	spiritSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(spiritPath)
	if err != nil {
		log.Fatalf("can not load spirit images file %s: %s", spiritPath, err)
	}
	spiritSpriteSheet := engine.NewSpriteSheet(spiritSpriteSheetImage, 1, 5, 50, 50)
	spiritSpriteSheet.SetFrameMap(map[string][]int{
		"idle": {0, 0, 50, 0, 100, 0, 150, 0, 200, 0},
	})
	spiritSpriteSheet.SetFrameSpeed(15)
	spirit := NewSpirit("spirit", spiritSpriteSheet, 30, 30)

	// male warrior
	warriorPath := filepath.Join(wd, "assets/images/male_warrior.png")
	warriorSpriteSheetImage, _, err := ebitenutil.NewImageFromFile(warriorPath)
	if err != nil {
		log.Fatalf("can not load male warrior images file %s: %s", warriorPath, err)
	}
	warriorSpriteSheet := engine.NewSpriteSheet(warriorSpriteSheetImage, 16, 4, 512, 512)
	warriorSpriteSheet.SetFrameMap(map[string][]int{
		"down":         {0, 512, 512, 512, 1024, 512, 1536, 512},
		"up":           {0, 2560, 512, 2560, 1024, 2560, 1536, 2560},
		"left":         {0, 4608, 512, 4608, 1024, 4608, 1536, 4608},
		"right":        {0, 6656, 512, 6656, 1024, 6656, 1536, 6656},
		"attack/down":  {0, 1024, 512, 1024, 1024, 1024, 1536, 1024},
		"attack/up":    {0, 3072, 512, 3072, 1024, 3072, 1536, 3072},
		"attack/left":  {0, 5120, 512, 5120, 1024, 5120, 1536, 5120},
		"attack/right": {0, 7168, 512, 7168, 1024, 7168, 1536, 7168},
	})
	warriorSpriteSheet.SetFrameSpeed(15)
	warrior := NewWarrior("warrior", warriorSpriteSheet, 0, 0)
	_ = warrior

	var menuitems []*engine.MenuItem = []*engine.MenuItem{
		engine.NewMenuItem("one"),
		engine.NewMenuItem("two"),
		engine.NewMenuItem("three"),
	}

	g := &Game{
		Tilemap: tilemap,
		Actors:  []engine.IActor{knight, spirit},
		//Actors: []engine.IActor{warrior},
		Camera: engine.NewCamera(0, 0, screenWidth, screenHeight),
		//warriorSpriteSheet: warriorSpriteSheet,
		menu: engine.NewSubMenu("menu", 10, 10, 100, 4, menuitems, 0, nil),
	}

	for _, act := range g.Actors {
		g.colliders = append(g.colliders, act)
	}

	event := NewEvent("event", 0, 60, 16, 16)
	g.colliders = append(g.colliders, event)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
