package gogame

import (
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// NewCanvas creates an empty canvas with the specified width and height.
func NewCanvas(width, height int) *Canvas {
	var err error
	canvas := &Canvas{
		rendererOutput: rendererOutput{
			textures: make(map[*sdl.Surface]*sdl.Texture),
			mask:     Color{1, 1, 1, 1},
		},
	}

	// no staticSurface flag, this suface is dynamic
	canvas.surface, err = sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0, 0, 0, 0)
	if err != nil {
		panic(errors.Wrap(err, "failed to create canvas"))
	}

	canvas.renderer, err = sdl.CreateSoftwareRenderer(canvas.surface)
	if err != nil {
		panic(errors.Wrap(err, "failed to create canvas"))
	}
	canvas.renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	return canvas
}

// Canvas is an offscreen picture that you can draw on.
type Canvas struct {
	surface *sdl.Surface
	rendererOutput
}

// Picture returns a pointer to the underlying picture of the canvas. The picture will change
// according to the drawing operations on the canvas.
func (c *Canvas) Picture() *Picture {
	return &Picture{
		surface: c.surface,
		rect:    sdl.Rect{X: 0, Y: 0, W: c.surface.W, H: c.surface.H},
	}
}

// OutputRect returns a (0, 0, w, h) rectangle, where w, h is the width and height of the canvas.
func (c *Canvas) OutputRect() Rect {
	return Rect{
		X: 0,
		Y: 0,
		W: float64(c.surface.W),
		H: float64(c.surface.H),
	}
}
