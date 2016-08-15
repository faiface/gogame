package gogame

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// Picture provies a basic interface for all kinds of software/hardware surfaces/textures.
type Picture interface {
	// Size returns the width and height of a picture.
	Size() (w, h int)

	// ColorAt returns the color of pixel (x, y) of a picture.
	ColorAt(x, y int) Color
}

// LoadPicture loads a picture stored in your filesystem at the specified path.
// If an error occured during the loading (e.g. file does not exist), an error will be returned.
func LoadPicture(path string) (Picture, error) {
	var (
		pic picture
		err error
	)
	pic.surface, err = img.Load(path)
	if err != nil {
		return nil, err
	}
	return pic, nil
}

type picture struct {
	surface *sdl.Surface
	texture *sdl.Texture
}

func (p picture) Size() (w, h int) {
	return int(p.surface.W), int(p.surface.H)
}

func (p picture) ColorAt(x, y int) Color {
	bpp := p.surface.BytesPerPixel()
	index := y*int(p.surface.W) + x
	pixel := p.surface.Pixels()[index*bpp : index*bpp+bpp]

	rgba := [4]float64{0.0, 0.0, 0.0, 1.0}
	for i := range pixel {
		rgba[i] = float64(pixel[i]) / 255
	}

	return Color{rgba[0], rgba[1], rgba[2], rgba[3]}
}
