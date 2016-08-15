package gogame

// Color is a RGBA representation of a color.
// All of the components should be between 0 and 1 (inclusive).
// If they are not, nevermind, I can deal with it.
type Color struct {
	R, G, B, A float64
}

// Colors defines some common colors.
var Colors = map[string]Color{
	"black":   {0.0, 0.0, 0.0, 1.0},
	"grey":    {0.5, 0.5, 0.5, 1.0},
	"white":   {1.0, 1.0, 1.0, 1.0},
	"red":     {1.0, 0.0, 0.0, 1.0},
	"green":   {0.0, 1.0, 0.0, 1.0},
	"blue":    {0.0, 0.0, 1.0, 1.0},
	"cyan":    {0.0, 1.0, 1.0, 1.0},
	"magenta": {1.0, 0.0, 1.0, 1.0},
	"yellow":  {1.0, 1.0, 0.0, 1.0},
}
