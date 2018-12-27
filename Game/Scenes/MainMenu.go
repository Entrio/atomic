package Scenes

import (
	"fmt"
	"github.com/Entrio/atomic/Game/LogManager"
	"github.com/Entrio/atomic/Game/ResourceManager"
	"time"
)

type MainMenuScene struct {
	Name            string
	ResourceManager *ResourceManager.ResourceManager
	Logger          *LogManager.LogManager
}

func (ms MainMenuScene) Awake() {
	if ms.Logger == nil {
		panic(fmt.Sprintf("Please pass log manager to scene %s", ms.Name))
	}

	if ms.ResourceManager == nil {
		panic(fmt.Sprintf("Please pass resource manager to scene %s", ms.Name))
	}

	ms.Logger.LogfTime("Main menu scene loaded")
}

func (ms MainMenuScene) Start() {
}

func (ms MainMenuScene) Draw2D() {
}

func (ms MainMenuScene) DrawUI() {
}

func (ms MainMenuScene) Draw3D() {
}

func (ms MainMenuScene) Update(time.Duration) {
}

func (ms MainMenuScene) GetName() string {
	return ms.Name
}
