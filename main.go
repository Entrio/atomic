package main

import (
	"github.com/Entrio/atomic/Game"
	"github.com/gen2brain/raylib-go/raylib"
	"os"
)

func main() {
	rl.InitWindow(800, 600, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)
	rl.HideCursor()

	/*
		camera := rl.Camera2D{
			Offset:   rl.Vector2{},
			Target:   rl.Vector2{},
			Rotation: 0,
			Zoom:     1.0,
		}
	*/

	g := Game.Newgame()
	//ent := Game.NewEntity(64, 64, 64, 64, g.LoadTexture("potions2"), rl.White)
	//ent2 := Game.NewEntity(128, 128, 32, 32, g.LoadTexture("potions"), rl.White)
	//r := float32(0.0)
	//r2 := float32(0.0)
	//mesh := rl.GenMeshPlane(32,32,20,20)
	//mesh.

	g.StartUpdateThread()
	for !g.ShouldExit() {
		g.Draw()
	}

	rl.CloseWindow()
	os.Exit(0)
}
