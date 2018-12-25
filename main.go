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

		/*
			if rl.IsKeyDown(rl.KeyQ) {
				camera.Rotation = camera.Rotation + 1.0
			}
			if rl.IsKeyDown(rl.KeyE) {
				camera.Rotation = camera.Rotation - 1.0
			}

			if rl.IsKeyDown(rl.KeyW) {
				camera.Offset = rl.Vector2{
					X: camera.Offset.X,
					Y: camera.Offset.Y + 10.0,
				}
			}

			if rl.IsKeyDown(rl.KeyS) {
				camera.Offset = rl.Vector2{
					X: camera.Offset.X,
					Y: camera.Offset.Y - 10.0,
				}
			}

			if rl.IsKeyDown(rl.KeyA) {
				camera.Offset = rl.Vector2{
					X: camera.Offset.X + 10.0,
					Y: camera.Offset.Y,
				}
			}

			if rl.IsKeyDown(rl.KeyD) {
				camera.Offset = rl.Vector2{
					X: camera.Offset.X - 10.0,
					Y: camera.Offset.Y,
				}
			}

			if rl.IsKeyDown(rl.KeyZ) {
				camera.Zoom = camera.Zoom + 0.2
			}

			if rl.IsKeyDown(rl.KeyX) {
				camera.Zoom = camera.Zoom - 0.2
			}
		*/
		/*
			r += 5.0
			r2 += 1.0
			if r > 360.0 {
				r = 0.0
			}
			if r2 > 360.0 {
				r2 = 0.0
			}

			ent.SetRotation(r)
			ent2.SetRotation(r2)
		*/

		//rl.BeginMode2D(camera)
		//go ent.Update(nil)
		//ent.Draw()

		//ent2.Draw()
		//rl.EndMode2D()

		g.Draw()
	}

	rl.CloseWindow()
	os.Exit(0)
}
