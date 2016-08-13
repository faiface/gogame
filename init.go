package gogame

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

// Init initializes Gogame (and SDL2). Call this before using Gogame.s
func Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return errors.New("failed to initialize SDL2")
	}
	return nil
}

// Quit deinitializes Gogame. Call this when you are done with Gogame.
func Quit() {
	sdl.Quit()
}
