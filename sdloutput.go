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
}

func newSdlOutput(window *sdl.Window, renderer *sdl.Renderer) *sdlOutput {
	return &sdlOutput{
		window:   window,
		renderer: renderer,
		textures: make(map[*sdl.Surface]*sdl.Texture),
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

func (o *sdlOutput) DrawLine(a, b Vec, thickness float64, color Color) {
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

func (o *sdlOutput) DrawPicture(x, y, w, h float64, pic *Picture) {
	if o.textures[pic.surface] == nil || pic.surface.Flags&staticSurface == 0 {
		var err error
		o.textures[pic.surface], err = o.renderer.CreateTextureFromSurface(pic.surface)
		if err != nil {
			panic("failed to create a texture from a surface")
		}
	}

	dst := sdl.Rect{
		X: int32(x + 0.5),
		Y: int32(y + 0.5),
		W: int32(w + 0.5),
		H: int32(h + 0.5),
	}
	o.renderer.Copy(o.textures[pic.surface], &pic.rect, &dst)
}
