package gogame

// This file internally implements output devices through SDL2.

import (
	"github.com/veandco/go-sdl2/sdl"
	gfx "github.com/veandco/go-sdl2/sdl_gfx"
)

type sdlOutput struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	textures map[*sdl.Surface]*sdl.Texture
	mask     Color
}

func newSdlOutput(window *sdl.Window, renderer *sdl.Renderer) *sdlOutput {
	return &sdlOutput{
		window:   window,
		renderer: renderer,
		textures: make(map[*sdl.Surface]*sdl.Texture),
		mask:     Color{1, 1, 1, 1},
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

func (o *sdlOutput) OutputRect() Rect {
	w, h := o.window.GetSize()
	return Rect{X: 0, Y: 0, W: float64(w), H: float64(h)}
}

func (o *sdlOutput) SetMask(color Color) {
	o.mask = color
}

func (o *sdlOutput) Clear(color Color) {
	color = color.Mul(o.mask)
	o.renderer.SetDrawColor(color.toSDLRGBA())
	o.renderer.Clear()
}

func (o *sdlOutput) DrawLine(a, b Vec, thickness float64, color Color) {
	color = color.Mul(o.mask)
	gfx.ThickLineColor(
		o.renderer,
		int(a.X+0.5),
		int(a.Y+0.5),
		int(b.X+0.5),
		int(b.Y+0.5),
		int(thickness+0.5),
		color.toSDL(),
	)
}

func (o *sdlOutput) DrawPolygon(points []Vec, thickness float64, color Color) {
	color = color.Mul(o.mask)
	if thickness == 0 {
		xInt16 := make([]int16, len(points))
		yInt16 := make([]int16, len(points))
		for i := 0; i < len(points); i++ {
			xInt16[i] = int16(points[i].X + 0.5)
			yInt16[i] = int16(points[i].Y + 0.5)
		}
		gfx.FilledPolygonColor(o.renderer, xInt16, yInt16, color.toSDL())
	} else {
		for i := 0; i < len(points); i++ {
			j := (i + 1) % len(points)
			x1, y1 := int(points[i].X+0.5), int(points[i].Y+0.5)
			x2, y2 := int(points[j].X+0.5), int(points[j].Y+0.5)
			gfx.ThickLineColor(o.renderer, x1, y1, x2, y2, int(thickness+0.5), color.toSDL())
			gfx.FilledCircleColor(o.renderer, x1, y1, int(thickness/2+0.5), color.toSDL())
		}
	}
}

func (o *sdlOutput) DrawRect(rect Rect, thickness float64, color Color) {
	color = color.Mul(o.mask)
	points := []Vec{
		{rect.X, rect.Y},
		{rect.X + rect.W, rect.Y},
		{rect.X + rect.W, rect.Y + rect.H},
		{rect.X, rect.Y + rect.H},
	}
	o.DrawPolygon(points, thickness, color)
}

func (o *sdlOutput) DrawPicture(rect Rect, pic *Picture) {
	if o.textures[pic.surface] == nil || pic.surface.Flags&staticSurface == 0 {
		var err error
		o.textures[pic.surface], err = o.renderer.CreateTextureFromSurface(pic.surface)
		if err != nil {
			panic("failed to create a texture from a surface")
		}
	}

	r, g, b, a := o.mask.toSDLRGBA()

	texture := o.textures[pic.surface]
	texture.SetColorMod(r, g, b)
	texture.SetAlphaMod(a)

	dst := sdl.Rect{
		X: int32(rect.X + 0.5),
		Y: int32(rect.Y + 0.5),
		W: int32(rect.W + 0.5),
		H: int32(rect.H + 0.5),
	}
	o.renderer.Copy(texture, &pic.rect, &dst)
}
