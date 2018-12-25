package Game

import (
	"github.com/Entrio/atomic/Game/LogManager"
	"github.com/Entrio/atomic/Game/ResourceManager"
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type (
	Game struct {
		LogManager      *LogManager.LogManager
		ResourceManager *ResourceManager.ResourceManager
		startTime       *time.Time
		shouldExit      bool
	}
)

// Create a new instance of the game
func Newgame() *Game {
	t := time.Now().Local()
	g := &Game{
		LogManager: LogManager.NewLogManager(&t),
		startTime:  &t,
	}

	g.ResourceManager = ResourceManager.NewResourceManager(g.LogManager)

	return g
}

func (g *Game) awake() {}
func (g *Game) start() {}

/**
Ok, we are ready to roll with the main update
*/
func (g *Game) StartUpdateThread() {
	// Do some work before update

	// Start the update goroutine
	go g.update()
}

func (g *Game) update() {
	for !g.shouldExit {
		//TODO: Update
		g.ProcessInput()
	}
}

func (g *Game) LoadTexture(name string) *rl.Texture2D {
	return g.ResourceManager.LoadTexture(name)
}

func (g *Game) ShouldExit() bool {
	return g.shouldExit
}
