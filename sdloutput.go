package gogame

import (
	"github.com/veandco/go-sdl2/sdl"
	gfx "github.com/veandco/go-sdl2/sdl_gfx"
)

type sdlOutput struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func newSdlOutput(window *sdl.Window, renderer *sdl.Renderer) *sdlOutput {
	return &sdlOutput{
		window:   window,
		renderer: renderer,
	}
}

func (o *sdlOutput) WindowSetTitle(title string) {
	o.window.SetTitle(title)
}

func (o *sdlOutput) WindowSetFullscreen(fullscreen bool) {
	var flags uint32
	if fullscreen {
		flags |= sdl.WINDOW_FULLSCREEN
	}
	o.window.SetFullscreen(flags)
}

func (o *sdlOutput) WindowResize(w, h int) {
	o.window.SetSize(w, h)
}

func (o *sdlOutput) Clear(color Color) {
	o.renderer.SetDrawColor(color.toSDLRGBA())
	o.renderer.Clear()
}

func (o *sdlOutput) DrawLine(x1, y1, x2, y2, thickness float64, color Color) {
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

func (o *sdlOutput) DrawPolygon(x, y []float64, thickness float64, color Color) {
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

func (o *sdlOutput) DrawPicture(x, y, w, h float64, pic Picture) {
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
