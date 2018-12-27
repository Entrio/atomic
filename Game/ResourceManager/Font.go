package ResourceManager

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"os"
)

func (rm *ResourceManager) LoadFont(name string) *rl.Font {
	if f, ok := rm.fonts[name]; ok {
		rm.logManager.Logf("Loaded font %s from memory", name)
		return &f
	}

	filename := fmt.Sprintf("%s/fonts/%s.png", *rm.GetConfig("base_path"), name)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		rm.logManager.LogfTime("Failed to load font %s. Error: %s", name, err.Error())
		return nil
	}

	f := rl.LoadFont(filename)
	rm.fonts[name] = f
	return nil
}
