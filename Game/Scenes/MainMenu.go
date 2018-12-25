package Scenes

import "time"

type MainMenuScene struct {
	Name string
}

func (ms MainMenuScene) GetName() string {
	return ms.Name
}

func (MainMenuScene) start() {
	panic("implement me")
}

func (MainMenuScene) awake() {
	panic("implement me")
}

func (MainMenuScene) drawUI() {
	panic("implement me")
}

func (MainMenuScene) draw2D() {
	panic("implement me")
}

func (MainMenuScene) draw3D() {
	panic("implement me")
}

func (MainMenuScene) update(time.Duration) {
	panic("implement me")
}
