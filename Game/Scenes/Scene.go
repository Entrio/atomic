package Scenes

import (
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type Scene interface {
	GetName() string
	Awake()
	Start()
	Draw2D()
	DrawUI()
	Draw3D()
	Update(time.Duration)
	GetCamera3D() *rl.Camera
}
