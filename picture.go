package gogame

// Picture provies a basic interface for all kinds of software/hardware surfaces/textures.
type Picture interface {
	// Size returns the width and height of a picture.
	Size() (w, h int)

	// ColorAt returns the color of pixel (x, y) of a picture.
	ColorAt(x, y int) Color
}
