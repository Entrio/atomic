package ResourceManager

import (
	"github.com/Entrio/atomic/Game/LogManager"
	"github.com/gen2brain/raylib-go/raylib"
	"os"
	"unsafe"
)

type (
	ResourceManager struct {
		textures      map[string]rl.Texture2D
		configuration map[string]interface{}
		logManager    *LogManager.LogManager
	}
)

// Initiate a new resource manager
func NewResourceManager(l *LogManager.LogManager) *ResourceManager {
	rm := &ResourceManager{
		textures:      map[string]rl.Texture2D{},
		configuration: map[string]interface{}{},
		logManager:    l,
	}

	path, _ := os.Getwd()
	rm.configuration["base_path"] = path
	l.Logf("RM > set base_path to %s", rm.configuration["base_path"])

	// Create a basic texture not found texture
	c := make([]rl.Color, 16*16*unsafe.Sizeof(rl.Color{}))
	defer func() {
		c = nil
	}()

	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			if ((x/4+y/4)/1)%2 == 0 {
				c[y*16+x] = rl.Pink
			} else {
				c[y*16+x] = rl.Green
			}
		}
	}

	rm.textures["not_found"] = rl.LoadTextureFromImage(rl.LoadImageEx(c, 16, 16))
	rl.SetTextureFilter(rm.textures["not_found"], rl.FilterPoint)

	l.LogfTime("Resource manager loaded.")

	return rm
}

// Attempt to load configuration variable
func (rm *ResourceManager) GetConfig(name string) *interface{} {
	if i, ok := rm.configuration[name]; ok {
		return &i
	}
	return nil
}
