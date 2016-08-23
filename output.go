package gogame

// Output combines all of the output methods together.
type Output interface {
	WindowOutput
	VideoOutput
	AudioOutput
}

// WindowOutput lets you resize a window, set it fullscreen or windowed.
type WindowOutput interface {
	// WindowSetTitle sets the title of the window to title.
	WindowSetTitle(title string)

	// WindowSetFullscreen sets window fullscreen if fullscreen is true. Otherwise
	// sets it windowed.
	WindowSetFullscreen(fullscreen bool)

	// WindowResize changes the size of a window (or resolution, if the window is fullscreen).
	WindowResize(w, h int)
}

// VideoOutput combines all of the video output methods together and adds a Clear method
// to clear the screen.
type VideoOutput interface {
	// OutputRect returns the output rectangle.
	OutputRect() Rect

	// Clear fill whole video output with one color.
	Clear(color Color)

	// SetMask sets a color with which every following draw call should be masked.
	// Masking means to multiply one color by another.
	// Default mask is Color{R: 1, G: 1, B: 1, A: 1}.
	SetMask(color Color)

	// DrawLine draws a line of specified thickness from point a to point b
	// using the provided color.
	DrawLine(a, b Vec, thickness float64, color Color)

	// DrawPolygon draws a closed polygon from the supplied points.
	// If the thickness is 0, the polygon will be filled.
	DrawPolygon(points []Vec, thickness float64, color Color)

	// DrawRect draws a rectangle parallel with the axis of the coordinate system.
	// If the thickness is 0, the rectangle will be filled.
	DrawRect(rect Rect, thickness float64, color Color)

	// DrawPicture draws a picture onto a rect. The picture will be
	// stretched to fit the rectangle.
	DrawPicture(rect Rect, pic *Picture)
}

// AudioOutput lets you play sounds and music.
type AudioOutput interface {
	//TODO
}
