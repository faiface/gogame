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
	PrimitiveVideoOutput
	PictureVideoOutput

	// Clear fill whole video output with one color.
	Clear(color Color)
}

// PrimitiveVideoOutput lets you draw primitive geometric shapes.
type PrimitiveVideoOutput interface {
	// DrawLine draws a line of specified thickness from point a to point b
	// using the provided color.
	DrawLine(a, b Vec, thickness float64, color Color)

	// DrawPolygon draws a closed polygon from the supplied points.
	// If the thickness is 0, the polygon will be filled.
	DrawPolygon(points []Vec, thickness float64, color Color)
}

// PictureVideoOutput lets you draw pictures.
type PictureVideoOutput interface {
	// DrawPicture draws a picture onto a rect. The picture will be
	// stretched to fit the rectangle.
	DrawPicture(rect Rect, pic *Picture)
}

// AudioOutput lets you play sounds and music.
type AudioOutput interface {
	//TODO
}
