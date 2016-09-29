package gogame

import "math"

// Vec represents a 2D vector.
type Vec struct {
	X, Y float64
}

// XY returns the componets of a vector in two return values.
func (u Vec) XY() (x, y float64) {
	return u.X, u.Y
}

// A adds two vectors.
func (u Vec) A(v Vec) Vec {
	return Vec{
		u.X + v.X,
		u.Y + v.Y,
	}
}

// S subtracts two vectors.
func (u Vec) S(v Vec) Vec {
	return Vec{
		u.X - v.X,
		u.Y - v.Y,
	}
}

// M multiplies a vector by a scalar.
func (u Vec) M(c float64) Vec {
	return Vec{
		u.X * c,
		u.Y * c,
	}
}

// D divides a vector by a scalar.
func (u Vec) D(c float64) Vec {
	return Vec{
		u.X / c,
		u.Y / c,
	}
}

// Len2 calculates a squared length of a vector.
func (u Vec) Len2() float64 {
	return u.X*u.X + u.Y*u.Y
}

// Len calculates a length of a vector.
func (u Vec) Len() float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y)
}

// Rect represents a rectangle in 2D space.
type Rect struct {
	X, Y, W, H float64
}

// XYWH returns the componets of a rectangle in four return values.
func (r Rect) XYWH() (x, y, w, h float64) {
	return r.X, r.Y, r.W, r.H
}

// Pos returns position of a rectangle as a vector.
func (r Rect) Pos() Vec {
	return Vec{r.X, r.Y}
}

// Size returns size of a rectangle as a vector.
func (r Rect) Size() Vec {
	return Vec{r.W, r.H}
}

// Center returns position of the center of a rectangle.
func (r Rect) Center() Vec {
	return Vec{r.X + r.W/2, r.Y + r.H/2}
}

// MovedTo returns a copy of a rectangle with position set to to pos.
func (r Rect) MovedTo(pos Vec) Rect {
	return Rect{
		X: pos.X,
		Y: pos.Y,
		W: r.W,
		H: r.H,
	}
}

// MovedBy retuns a copy of a rectangle relatively moved by delta.
func (r Rect) MovedBy(delta Vec) Rect {
	return Rect{
		X: r.X + delta.X,
		Y: r.Y + delta.Y,
		W: r.W,
		H: r.H,
	}
}

// Resized returns a copy of a recangle with a new size.
func (r Rect) Resized(size Vec) Rect {
	return Rect{
		X: r.X,
		Y: r.Y,
		W: size.X,
		H: size.Y,
	}
}

// Overlap returns the overlap vector of r1 and r2. If there's no overlap the result is {0, 0}.
// Otherwise, the result represents, how much r1 needs to be moved not to overlap with r2.
// Rectangle r1 only needs to be moved by one component of result.
func Overlap(r1, r2 Rect) Vec {
	left := (r2.X + r2.W) - r1.X
	right := r2.X - (r1.X + r1.W)
	bottom := (r2.Y + r2.H) - r1.Y
	top := r2.Y - (r1.Y + r1.H)

	if left <= 0 || right >= 0 || bottom <= 0 || top >= 0 {
		return Vec{0, 0}
	}

	overlap := Vec{left, bottom}
	if math.Abs(right) < math.Abs(left) {
		overlap.X = right
	}
	if math.Abs(top) < math.Abs(bottom) {
		overlap.Y = top
	}

	return overlap
}
