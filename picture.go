package gogame

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// LoadPicture loads a picture from a file stored at the specified path.
// If the loading fails, an error is returned.
func LoadPicture(path string) (*Picture, error) {
	var (
		pic Picture
		err error
	)
	pic.surface, err = img.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load picture: %s", path)
	}
	pic.surface.Flags |= staticSurface
	pic.rect = sdl.Rect{X: 0, Y: 0, W: pic.surface.W, H: pic.surface.H}
	return &pic, nil
}

// Picture is a static raster image, usually loaded from a file.
type Picture struct {
	surface *sdl.Surface
	rect    sdl.Rect
}

// Size returns the width and height of a picture in pixels.
func (p *Picture) Size() (w, h int) {
	return int(p.surface.W), int(p.surface.H)
}

// Slice cuts a rectangle (x, y, w, h) from a picture.
func (p *Picture) Slice(x, y, w, h int) *Picture {
	return &Picture{
		surface: p.surface,
		rect: sdl.Rect{
			X: p.rect.X + int32(x),
			Y: p.rect.Y + int32(y),
			W: int32(w),
			H: int32(h),
		},
	}
}

// Copy creates an exact independent copy of a picture.
// This is particularly useful when dealing with canvases, since this copy can be rendered
// more effeciently than the internal picture of a canvas.
func (p *Picture) Copy() *Picture {
	surface, err := sdl.CreateRGBSurface(
		p.surface.Flags,
		p.surface.W,
		p.surface.H,
		int32(p.surface.Format.BitsPerPixel),
		p.surface.Format.Rmask,
		p.surface.Format.Gmask,
		p.surface.Format.Bmask,
		p.surface.Format.Amask,
	)

	if err != nil {
		panic(fmt.Errorf("failed to copy picture: %s", err))
	}

	p.surface.Blit(nil, surface, nil)
	surface.Flags |= staticSurface

	return &Picture{
		surface: surface,
		rect:    p.rect,
	}
}

const (
	staticSurface = 1 << iota
)
