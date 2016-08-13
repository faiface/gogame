package gogame

// Color is a RGBA representation of a color.
// All of the components should be between 0 and 1 (inclusive).
// If they are not, nevermind, I can deal with it.
type Color struct {
	R, G, B, A float64
}
