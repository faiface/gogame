package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// WindowFlag is a type of window flags.
type WindowFlag int

const (
	// WindowFullscreen makes window fullscreen.
	WindowFullscreen WindowFlag = iota

	// WindowBorderless removes window decoration.
	WindowBorderless

	// WindowResizable makes window resizable.
	WindowResizable
)

// Window is a handle to a window and its canvas.
type Window struct {
	sdlWindow *sdl.Window
}

// MakeWindow creates a new window with given parameters. If the window creation fails, this
// function panics. If you want to handle this failure instead, use MakeWindowErr.
func MakeWindow(title string, width, height int, flags ...WindowFlag) *Window {
	window, err := MakeWindowErr(title, width, height, flags...)
	if err != nil {
		panic(fmt.Sprintf("Gogame: MakeWindow failed: %s", err))
	}
	return window
}

// MakeWindowErr creates a new window with given parameters. If the window creation fails, this
// function returns an error.
func MakeWindowErr(title string, width, height int, flags ...WindowFlag) (*Window, error) {
	var sdlFlags uint32
	for _, flag := range flags {
		switch flag {
		case WindowFullscreen:
			sdlFlags |= sdl.WINDOW_FULLSCREEN_DESKTOP
		case WindowBorderless:
			sdlFlags |= sdl.WINDOW_BORDERLESS
		case WindowResizable:
			sdlFlags |= sdl.WINDOW_RESIZABLE
		}
	}

	sdlWindow, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width,
		height,
		sdlFlags,
	)

	return &Window{sdlWindow: sdlWindow}, err
}

// Close closes and destroys a window. Don't use the window after it's been closed.
func (w *Window) Close() {
	w.sdlWindow.Destroy()
}
