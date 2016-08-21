package gogame

// This file internally implements input interfaces through SDL2.

import "github.com/veandco/go-sdl2/sdl"

type sdlInput struct {
	window                                 *sdl.Window
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

func newSdlInput(window *sdl.Window) *sdlInput {
	input := sdlInput{
		window:       window,
		prevMouse:    make(map[int]bool),
		mouse:        make(map[int]bool),
		prevKeyboard: make(map[int]bool),
		keyboard:     make(map[int]bool),
	}

	input.windowX, input.windowY = window.GetPosition()
	input.windowW, input.windowH = window.GetSize()
	input.windowHasFocus = window.GetFlags()&sdl.WINDOW_INPUT_FOCUS != 0
	input.windowGainedFocus = input.windowHasFocus
	input.mouseX, input.mouseY, _ = sdl.GetMouseState()
	input.prevMouseX, input.prevMouseY = input.mouseX, input.mouseY

	return &input
}

func (i *sdlInput) WindowPosition() (x, y int) { return i.windowX, i.windowY }
func (i *sdlInput) WindowSize() (w, h int)     { return i.windowW, i.windowH }
func (i *sdlInput) WindowMoved() bool          { return i.windowMoved }
func (i *sdlInput) WindowResized() bool        { return i.windowResized }
func (i *sdlInput) WindowClosed() bool         { return i.windowClosed }
func (i *sdlInput) WindowHasFocus() bool       { return i.windowHasFocus }
func (i *sdlInput) WindowLostFocus() bool      { return i.windowLostFocus }
func (i *sdlInput) WindowGainedFocus() bool    { return i.windowGainedFocus }

func (i *sdlInput) MousePosition() (x, y int)     { return i.mouseX, i.mouseY }
func (i *sdlInput) MouseDelta() (dx, dy int)      { return i.mouseX - i.prevMouseX, i.mouseY - i.prevMouseY }
func (i *sdlInput) MouseDown(button int) bool     { return i.mouse[button] }
func (i *sdlInput) MouseJustDown(button int) bool { return i.mouse[button] && !i.prevMouse[button] }
func (i *sdlInput) MouseJustUp(button int) bool   { return !i.mouse[button] && i.prevMouse[button] }

func (i *sdlInput) KeyDown(key int) bool     { return i.keyboard[key] }
func (i *sdlInput) KeyJustDown(key int) bool { return i.keyboard[key] && !i.prevKeyboard[key] }
func (i *sdlInput) KeyJustUp(key int) bool   { return !i.keyboard[key] && i.prevKeyboard[key] }

func (i *sdlInput) update() {
	i.windowMoved = false
	i.windowResized = false
	i.windowClosed = false
	i.windowHasFocus = i.window.GetFlags()&sdl.WINDOW_INPUT_FOCUS != 0
	i.windowGainedFocus = false
	i.windowLostFocus = false

	for button := range i.mouse {
		i.prevMouse[button] = i.mouse[button]
	}
	for key := range i.keyboard {
		i.prevKeyboard[key] = i.keyboard[key]
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.QuitEvent:
			i.windowClosed = true
		case *sdl.WindowEvent:
			switch event.Event {
			case sdl.WINDOWEVENT_MOVED:
				i.windowMoved = true
			case sdl.WINDOWEVENT_RESIZED, sdl.WINDOWEVENT_SIZE_CHANGED:
				i.windowResized = true
			case sdl.WINDOWEVENT_FOCUS_GAINED:
				i.windowGainedFocus = true
			case sdl.WINDOWEVENT_FOCUS_LOST:
				i.windowLostFocus = true
			case sdl.WINDOWEVENT_CLOSE:
				i.windowClosed = true
			}
		case *sdl.MouseButtonEvent:
			switch event.Type {
			case sdl.MOUSEBUTTONDOWN:
				i.mouse[int(event.Button)] = true
			case sdl.MOUSEBUTTONUP:
				i.mouse[int(event.Button)] = false
			}
		case *sdl.KeyDownEvent:
			i.keyboard[int(event.Keysym.Sym)] = true
		case *sdl.KeyUpEvent:
			i.keyboard[int(event.Keysym.Sym)] = false
		}
	}

	i.windowX, i.windowY = i.window.GetPosition()
	i.windowW, i.windowH = i.window.GetSize()
	i.prevMouseX, i.prevMouseY = i.mouseX, i.mouseY
	i.mouseX, i.mouseY, _ = sdl.GetMouseState()
}
