package gogame

import (
	"errors"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
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

	outputHandler := output{
		window:   window,
		renderer: renderer,
	}

	timer := time.Now()

	var framerate <-chan time.Time
	if !cfg.VSync && cfg.FPS != 0 {
		framerate = time.Tick(time.Second / time.Duration(cfg.FPS))
	}

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

		dt := float64(timer.Sub(time.Now())) / float64(time.Second)
		timer = time.Now()

		lf(Context{
			Dt:     dt,
			Input:  &inputState,
			Output: &outputHandler,
		})

		if framerate != nil {
			<-framerate
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

type output struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func (o *output) WindowSetFullscreen(fullscreen bool) {
	var flags uint32
	if fullscreen {
		flags |= sdl.WINDOW_FULLSCREEN
	}
	o.window.SetFullscreen(flags)
}

func (o *output) WindowResize(w, h int) {
	o.window.SetSize(w, h)
}

func (o *output) Clear(color Color) {
	o.renderer.SetDrawColor(color.toSDLRGBA())
	o.renderer.Clear()
}

func (o *output) DrawLine(x1, y1, x2, y2, thickness float64, color Color) {
	gfx.ThickLineColor(
		o.renderer,
		int(x1+0.5),
		int(y1+0.5),
		int(x2+0.5),
		int(y2+0.5),
		int(thickness+0.5),
		color.toSDL().Uint32(),
	)
}

func (o *output) DrawPolygon(x, y []float64, thickness float64, color Color) {
	var numPoints int
	if len(x) < len(y) {
		numPoints = len(x)
	} else {
		numPoints = len(y)
	}

	if thickness == 0 {
		xInt16 := make([]int16, numPoints)
		yInt16 := make([]int16, numPoints)
		for i := 0; i < numPoints; i++ {
			xInt16[i] = int16(x[i] + 0.5)
			yInt16[i] = int16(y[i] + 0.5)
		}
		gfx.FilledPolygonColor(o.renderer, xInt16, yInt16, color.toSDL().Uint32())
	} else {
		for i := 0; i < numPoints; i++ {
			j := (i + 1) % numPoints
			x1, y1 := int(x[i]+0.5), int(y[i]+0.5)
			x2, y2 := int(x[j]+0.5), int(y[j]+0.5)
			gfx.ThickLineColor(o.renderer, x1, y1, x2, y2, int(thickness+0.5), color.toSDL().Uint32())
		}
	}
}

func (o *output) DrawPicture(x, y, w, h float64, pic Picture) {
	if pic, ok := pic.(picture); ok {
		if pic.texture == nil {
			var err error
			pic.texture, err = o.renderer.CreateTextureFromSurface(pic.surface)
			if err != nil {
				panic("creating texture failed")
			}
		}
		o.renderer.Copy(pic.texture, nil, &sdl.Rect{
			X: int32(x + 0.5),
			Y: int32(y + 0.5),
			W: int32(w + 0.5),
			H: int32(h + 0.5),
		})
	}
}
