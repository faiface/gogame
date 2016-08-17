package gogame

// This file contains internal functions to create SDL2 windows and renderers.

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

func makeWindow(cfg Config) (*sdl.Window, error) {
	var winFlags uint32
	if cfg.Resizable {
		winFlags |= sdl.WINDOW_RESIZABLE
	}
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
		return nil, errors.New("failed to create a window")
	}

	return window, nil
}

func makeRenderer(cfg Config, window *sdl.Window) (*sdl.Renderer, error) {
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
		return nil, errors.New("failed to create a renderer")
	}

	return renderer, nil
}
