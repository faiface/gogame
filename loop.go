package gogame

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

// LoopFunc is a type of a game loop body function.
type LoopFunc func(ctx Context)

// Loop opens a game window based on the provided config and starts a game loop.
// At each iteration of the game loop, it calls the provided LoopFunc.
func Loop(cfg Config, lf LoopFunc) error {
	var winFlags uint32
	if cfg.Fullscreen {
		winFlags |= sdl.WINDOW_FULLSCREEN
	}
	if cfg.Borderless {
		winFlags |= sdl.WINDOW_BORDERLESS
	}

	window, err := sdl.CreateWindow(
		cfg.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		cfg.Width,
		cfg.Height,
		winFlags,
	)
	if err != nil {
		return errors.New("failed to create a window")
	}
	defer window.Destroy()

	var rendFlags uint32
	if cfg.SoftwareRender {
		rendFlags |= sdl.RENDERER_SOFTWARE
	} else {
		rendFlags |= sdl.RENDERER_ACCELERATED
	}
	if cfg.VSync {
		rendFlags |= sdl.RENDERER_PRESENTVSYNC
	}

	renderer, err := sdl.CreateRenderer(window, 0, rendFlags)
	if err != nil {
		return errors.New("failed to create a renderer")
	}
	defer renderer.Destroy()

	return nil
}
