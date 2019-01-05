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
}

var (
	plane               *rl.Model
	cam                 *rl.Camera3D
	winHeight, winWidth int
)

func (ms MainMenuScene) GetCamera3D() *rl.Camera3D {
	return cam
}

func (ms MainMenuScene) Awake() {
	if ms.Logger == nil {
		panic(fmt.Sprintf("Please pass log manager to scene %s", ms.Name))
	}

	if ms.ResourceManager == nil {
		panic(fmt.Sprintf("Please pass resource manager to scene %s", ms.Name))
	}

	winWidth = rl.GetScreenWidth()
	winHeight = rl.GetScreenHeight()

	ms.Logger.LogfTime("Main menu scene loaded")

	cam = &rl.Camera3D{
		Position: rl.NewVector3(0.0, 5.0, 5.0),
		Target:   rl.NewVector3(-5, 0.0, 0.0),
		Up:       rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:     60.0,
		Type:     rl.CameraPerspective,
	}

	ms.Logger.LogfTime("Main menu has awoken!")
}

func (ms MainMenuScene) Start() {
	rl.SetCameraMode(*cam, rl.CameraFree)

	p := rl.Mesh{
		VertexCount:   0,
		TriangleCount: 0,
		Vertices:      nil,
		Texcoords:     nil,
		Texcoords2:    nil,
		Normals:       nil,
		Tangents:      nil,
		Colors:        nil,
		Indices:       nil,
		VaoID:         0,
		VboID:         [7]uint32{},
	}
	_ = p
	//v := []float32{
	//	1, 2, 3,
	//}
	//c := []uint8{255, 230, 250}
	//m := rl.NewMesh(3, 1, &v, nil, nil, nil, nil, &c, nil, 1, [7]uint32{})
	m := rl.LoadModelFromMesh(rl.GenMeshPlane(5, 5, 2, 2))
	plane = &m
	plane.Material.Maps[rl.MapDiffuse].Texture = rl.LoadTextureFromImage(rl.GenImageChecked(2, 2, 1, 1, rl.Red, rl.Green))
}

func (ms MainMenuScene) Draw2D() {
}

func (ms MainMenuScene) DrawUI() {
	rl.DrawRectangle(5, int32(winHeight-35), 310, 30, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(5, int32(winHeight-35), 310, 30, rl.Fade(rl.DarkBlue, 0.5))
	rl.DrawText(fmt.Sprintf(
		"Position: X:%f Y:%f Z:%f | %d",
		cam.Position.X,
		cam.Position.Y,
		cam.Position.Z,
		plane.Mesh.VertexCount,
	), 10, int32(winHeight-25), 10, rl.Blue)
}

func (ms MainMenuScene) Draw3D() {
	rl.BeginMode3D(*cam)
	rl.DrawModel(*plane, rl.Vector3{}, 1, rl.White)
	rl.DrawGrid(10, 1)
	rl.EndMode3D()
}

func (ms MainMenuScene) Update(dt time.Duration) {

}

func (ms MainMenuScene) GetName() string {
	return ms.Name
}
