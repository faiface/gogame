package gogame

// Camera allows moving and zooming the game view.
// It achieves so by transfroming points between so called 'game space' and 'display space'.
// Game space represents coordinates used internally inside a game.
// Display space represents coordinates of pixels on the screen (where (0, 0) is a corner).
type Camera struct {
	CornerX, CornerY float64
	ZoomX, ZoomY     float64

	// VideoOutput is used to actually draw using a camera.
	VideoOutput
}

// Project transfroms a point from game space to display space.
func (c *Camera) Project(x, y float64) (float64, float64) {
	return (x - c.CornerX) * c.ZoomX, (y - c.CornerY) * c.ZoomY
}

// Unproject transfroms a point from display space to game space.
func (c *Camera) Unproject(x, y float64) (float64, float64) {
	return x/c.ZoomX + c.CornerX, y/c.ZoomY + c.CornerY
}

// ProjectVec transforms a vector from game space to display space.
func (c *Camera) ProjectVec(u Vec) (v Vec) {
	v.X, v.Y = c.Project(u.X, u.Y)
	return
}

// UnprojectVec transforms a vector from display space to game space.
func (c *Camera) UnprojectVec(u Vec) (v Vec) {
	v.X, v.Y = c.Unproject(u.X, u.Y)
	return
}

// ProjectRect transforms a rectangle from game space to display space.
func (c *Camera) ProjectRect(r1 Rect) (r2 Rect) {
	if c.ZoomX < 0 {
		r1.X += r1.W
		r1.W *= -1
	}
	if c.ZoomY < 0 {
		r1.Y += r1.H
		r1.H *= -1
	}
	r2.X, r2.Y = c.Project(r1.X, r1.Y)
	r2.W, r2.H = r1.W*c.ZoomX, r1.H*c.ZoomY
	return
}

// UnprojectRect transfroms a rectangle from display space to game space.
func (c *Camera) UnprojectRect(r1 Rect) (r2 Rect) {
	if c.ZoomX < 0 {
		r1.X += r1.W
		r1.W *= -1
	}
	if c.ZoomY < 0 {
		r1.Y += r1.H
		r1.H *= -1
	}
	r2.X, r2.Y = c.Unproject(r1.X, r1.Y)
	r2.W, r2.H = r1.W/c.ZoomX, r1.H/c.ZoomY
	return
}

// OutputRect returns an unprojected output rectangle of the underlying video output.
func (c *Camera) OutputRect() Rect {
	return c.UnprojectRect(c.VideoOutput.OutputRect())
}

// DrawPoint draws a point projected with the camera using the underlying video output.
func (c *Camera) DrawPoint(point Vec, color Color) {
	point = c.ProjectVec(point)
	c.VideoOutput.DrawPoint(point, color)
}

// DrawLine draws a line projected with the camera using the underlying video output.
func (c *Camera) DrawLine(a, b Vec, thickness float64, color Color) {
	a, b = c.ProjectVec(a), c.ProjectVec(b)
	c.VideoOutput.DrawLine(a, b, thickness, color)
}

// DrawPolygon draws a polygon projected with the camera using the underlying video output.
func (c *Camera) DrawPolygon(points []Vec, thickness float64, color Color) {
	for i := range points {
		points[i] = c.ProjectVec(points[i])
	}
	c.VideoOutput.DrawPolygon(points, thickness, color)
}

// DrawRect draws a rectangle projected with the camera using the underlying video output.
func (c *Camera) DrawRect(rect Rect, thickness float64, color Color) {
	rect = c.ProjectRect(rect)
	c.VideoOutput.DrawRect(rect, thickness, color)
}

// DrawPicture draws a picture projected with the camera using the underlying video output.
func (c *Camera) DrawPicture(rect Rect, pic *Picture) {
	rect = c.ProjectRect(rect)
	c.VideoOutput.DrawPicture(rect, pic)
}
