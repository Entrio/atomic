package Game

import "github.com/gen2brain/raylib-go/raylib"

func (g *Game) ProcessInput() {
	if rl.IsKeyDown(rl.KeyEscape) {

		g.shouldExit = true
	}
}
