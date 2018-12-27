package Game

import (
	"github.com/Entrio/atomic/Game/LogManager"
	"github.com/Entrio/atomic/Game/ResourceManager"
	"github.com/Entrio/atomic/Game/Scenes"
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type (
	Game struct {
		LogManager      *LogManager.LogManager
		ResourceManager *ResourceManager.ResourceManager
		startTime       *time.Time
		scenes          map[string]Scenes.Scene
		activeScene     *Scenes.Scene
		shouldExit      bool
	}
)

// Create a new instance of the game
func Newgame() *Game {
	t := time.Now().Local()
	g := &Game{
		LogManager:  LogManager.NewLogManager(&t),
		startTime:   &t,
		activeScene: new(Scenes.Scene),
	}

	g.ResourceManager = ResourceManager.NewResourceManager(g.LogManager)

	scenes := make(map[string]Scenes.Scene)

	menuScene := Scenes.MainMenuScene{
		Name:            "main_menu",
		ResourceManager: g.ResourceManager,
		Logger:          g.LogManager,
	}

	g.LogManager.LogfTime("Attempting to load main menu scene")

	scenes[menuScene.GetName()] = menuScene

	*g.activeScene = scenes[menuScene.GetName()]

	g.scenes = scenes

	g.awake()

	return g
}

/**
Awake is called when all components are initialized and the new game instance is about to be returned
*/
func (g *Game) awake() {
	(*g.activeScene).Awake()
}

/**
Start is called *JUST* before the update loop enters its cycle
*/
func (g *Game) start() {
	(*g.activeScene).Start()
}

/**
Ok, we are ready to roll with the main update
*/
func (g *Game) StartUpdateThread() {
	// Do some work before update
	g.start()

	// Start the update goroutine
	go g.update()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Color{
		R: 100,
		G: 149,
		B: 237,
		A: 255,
	})

	g.Draw3D()
	g.Draw2D()
	g.DrawUI()

	rl.EndDrawing()
}

func (g *Game) LoadTexture(name string) *rl.Texture2D {
	return g.ResourceManager.LoadTexture(name)
}

func (g *Game) ShouldExit() bool {
	return g.shouldExit
}
