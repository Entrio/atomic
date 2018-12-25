package ResourceManager

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"os"
)

// Attempt to load the textured from the texture map, if it is not there load it from  file
func (rm *ResourceManager) LoadTexture(name string) *rl.Texture2D {
	if t, ok := rm.textures[name]; ok {
		rm.logManager.Logf("Loaded %s from memory", name)
		return &t
	}

	filename := fmt.Sprintf("%s/textures/%s.png", *rm.GetConfig("base_path"), name)

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		i := rm.textures["not_found"]
		return &i
	}

	rm.logManager.Logf("Attempting to load texture: %s", filename)
	// this texture is not found in the texture map, load it from file
	tex := rl.LoadTexture(fmt.Sprintf("%s/textures/%s.png", *rm.GetConfig("base_path"), name))

	if tex.ID == 0 && tex.Format == 0 && tex.Height == 0 && tex.Mipmaps == 0 && tex.Width == 0 {
		// failed to load texture!
		i := rm.textures["not_found"]
		return &i
	}

	rm.textures[name] = tex
	t := rm.textures[name]

	return &t
}
