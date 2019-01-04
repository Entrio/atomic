package Scenes

import (
	"fmt"
	"github.com/Entrio/atomic/Game/LogManager"
	"github.com/Entrio/atomic/Game/ResourceManager"
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type MainMenuScene struct {
	Name            string
	ResourceManager *ResourceManager.ResourceManager
	Logger          *LogManager.LogManager
	Camera          rl.Camera3D
}

var (
	plane rl.Model
)

func (ms MainMenuScene) GetCamera3D() *rl.Camera3D {
	return &ms.Camera
}

func (ms MainMenuScene) Awake() {
	if ms.Logger == nil {
		panic(fmt.Sprintf("Please pass log manager to scene %s", ms.Name))
	}

	if ms.ResourceManager == nil {
		panic(fmt.Sprintf("Please pass resource manager to scene %s", ms.Name))
	}

	//TODO: Load any resources that are needed

	ms.Logger.LogfTime("Main menu scene loaded")

	ms.Camera = rl.Camera3D{
		Position: rl.Vector3{X: 5.0, Y: 5.0, Z: 5.0},
		Target:   rl.Vector3{},
		Up:       rl.Vector3{Y: 1.0},
		Fovy:     45.0,
	}

	plane = rl.LoadModelFromMesh(rl.GenMeshPlane(2, 2, 5, 5))
	plane.Material.Maps[rl.MapDiffuse].Texture = *ms.ResourceManager.LoadTexture("fake")
	ms.Logger.LogfTime("Mainmenu has awoken!")
}

func (ms MainMenuScene) Start() {
	rl.SetCameraMode(ms.Camera, rl.CameraOrbital)
}

func (ms MainMenuScene) Draw2D() {
}

func (ms MainMenuScene) DrawUI() {
	rl.DrawRectangle(30, 400, 310, 30, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(30, 400, 310, 30, rl.Fade(rl.DarkBlue, 0.5))
	rl.DrawText("MOUSE LEFT BUTTON to CYCLE PROCEDURAL MODELS", 40, 410, 10, rl.Blue)
}

func (ms MainMenuScene) Draw3D() {

	rl.BeginMode3D(ms.Camera)
	rl.DrawModel(plane, rl.Vector3{}, 1, rl.White)
	rl.DrawGrid(10, 1)
	rl.EndMode3D()
}

func (ms MainMenuScene) Update(dt time.Duration) {

}

func (ms MainMenuScene) GetName() string {
	return ms.Name
}
