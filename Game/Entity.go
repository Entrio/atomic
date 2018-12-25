package Game

import (
	"github.com/gen2brain/raylib-go/raylib"
	"time"
)

type Entity struct {
	texture   *rl.Texture2D
	position  rl.Vector2
	rotation  float32
	rectangle rl.Rectangle
	color     rl.Color
}

func (e *Entity) SetRotation(r float32) {
	e.rotation = r
}

func NewEntity(x, y float32, width, height float32, tex *rl.Texture2D, c rl.Color) Entity {
	return Entity{
		texture: tex,
		position: rl.Vector2{
			X: x,
			Y: y,
		},
		rotation: 0.0,
		rectangle: rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  width,
			Height: height,
		},
		color: c,
	}
}

/**
Update anything that is required
 */
func (e *Entity) Update(d *time.Duration) {
	time.Sleep(time.Second * 1)
}

/**
Draw this entity to screen
 */
func (e *Entity) Draw() {
	rl.DrawTexturePro(
		*e.texture,
		rl.Rectangle{
			X:      64,
			Y:      0,
			Width:  16,
			Height: 16,
		},
		rl.Rectangle{
			X:      e.position.X,
			Y:      e.position.Y,
			Width:  float32(e.rectangle.Width),
			Height: float32(e.rectangle.Height),
		},
		rl.Vector2{
			X: e.rectangle.Width - (e.rectangle.Width / 2),
			Y: e.rectangle.Height - (e.rectangle.Height / 2),
		},
		e.rotation,
		e.color,
	)
}
