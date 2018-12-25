package Game

import "time"

type Scene interface {
	GetName() string
	start()
	awake()
	draw2D()
	drawUI()
	draw3D()
	update(time.Duration)
}
