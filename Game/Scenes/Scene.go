package Scenes

import (
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
}
