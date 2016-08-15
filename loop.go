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

	var inputState input
	inputState.windowX, inputState.windowY = window.GetPosition()
	inputState.windowW, inputState.windowH = window.GetSize()
	inputState.windowHasFocus = window.GetFlags()&sdl.WINDOW_INPUT_FOCUS != 0
	inputState.windowGainedFocus = inputState.windowHasFocus
	inputState.mouseX, inputState.mouseY, _ = sdl.GetMouseState()
	inputState.prevMouseX, inputState.prevMouseY = inputState.mouseX, inputState.mouseY

	for {
		inputState.windowMoved = false
		inputState.windowResized = false
		inputState.windowClosed = false
		inputState.windowHasFocus = window.GetFlags()&sdl.WINDOW_INPUT_FOCUS != 0
		inputState.windowGainedFocus = false
		inputState.windowLostFocus = false

		inputState.prevMouseX, inputState.prevMouseY = inputState.mouseX, inputState.mouseY
		inputState.mouseX, inputState.mouseY, _ = sdl.GetMouseState()

		inputState.prevMouse, inputState.mouse = inputState.mouse, inputState.prevMouse
		inputState.prevKeyboard, inputState.keyboard = inputState.keyboard, inputState.prevKeyboard

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				inputState.windowClosed = true
			case *sdl.WindowEvent:
				switch event.Event {
				case sdl.WINDOWEVENT_MOVED:
					inputState.windowMoved = true
				case sdl.WINDOWEVENT_RESIZED, sdl.WINDOWEVENT_SIZE_CHANGED:
					inputState.windowResized = true
				case sdl.WINDOWEVENT_FOCUS_GAINED:
					inputState.windowGainedFocus = true
				case sdl.WINDOWEVENT_FOCUS_LOST:
					inputState.windowLostFocus = true
				case sdl.WINDOWEVENT_CLOSE:
					inputState.windowClosed = true
				}
			case *sdl.MouseButtonEvent:
				switch event.Type {
				case sdl.MOUSEBUTTONDOWN:
					inputState.mouse[int(event.Button)] = true
				case sdl.MOUSEBUTTONUP:
					inputState.mouse[int(event.Button)] = false
				}
			case *sdl.KeyDownEvent:
				inputState.keyboard[int(event.Keysym.Sym)] = true
			case *sdl.KeyUpEvent:
				inputState.keyboard[int(event.Keysym.Sym)] = false
			}
		}
	}
}

type input struct {
	windowX, windowY, windowW, windowH     int
	windowMoved                            bool
	windowResized                          bool
	windowClosed                           bool
	windowHasFocus                         bool
	windowLostFocus                        bool
	windowGainedFocus                      bool
	prevMouseX, prevMouseY, mouseX, mouseY int
	prevMouse, mouse                       map[int]bool
	prevKeyboard, keyboard                 map[int]bool
}

func (i *input) WindowPosition() (x, y int) { return i.windowX, i.windowY }
func (i *input) WindowSize() (w, h int)     { return i.windowW, i.windowH }
func (i *input) WindowMoved() bool          { return i.windowMoved }
func (i *input) WindowResized() bool        { return i.windowResized }
func (i *input) WindowClosed() bool         { return i.windowClosed }
func (i *input) WindowHasFocus() bool       { return i.windowHasFocus }
func (i *input) WindowLostFocus() bool      { return i.windowLostFocus }
func (i *input) WindowGainedFocus() bool    { return i.windowGainedFocus }

func (i *input) MousePosition() (x, y int)     { return i.mouseX, i.mouseY }
func (i *input) MouseDelta() (dx, dy int)      { return i.mouseX - i.prevMouseX, i.mouseY - i.prevMouseY }
func (i *input) MouseDown(button int) bool     { return i.mouse[button] }
func (i *input) MouseJustDown(button int) bool { return i.mouse[button] && !i.prevMouse[button] }

func (i *input) KeyDown(key int) bool     { return i.keyboard[key] }
func (i *input) KeyJustDown(key int) bool { return i.keyboard[key] && !i.prevKeyboard[key] }
