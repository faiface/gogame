package gogame

// NewAnimation creates a non-cyclic animation with a specified duration.
func NewAnimation(frames []*Picture, duration float64) *Animation {
	return &Animation{
		Frames:   frames,
		Duration: duration,
		Cycle:    false,
	}
}

// NewCyclicAnimation creates a cyclic (repeating) animation with a specified duration of one cycle.
func NewCyclicAnimation(frames []*Picture, duration float64) *Animation {
	return &Animation{
		Frames:   frames,
		Duration: duration,
		Cycle:    true,
	}
}

// Animation is a sequence of frames that can be shown one after another to create an illusion of
// motion. It takes Duration time (in seconds) to play the whole animation.
// If Cycle is set to true, animation will be repeated over and over again.
type Animation struct {
	Frames   []*Picture
	Duration float64
	Cycle    bool
}

// FrameAt returns a frame at the specified time from the beginning of the animation.
// If the time is 0, the first frame is returned and so on.
func (a *Animation) FrameAt(time float64) *Picture {
	index := int(time / (a.Duration / float64(len(a.Frames))))

	if a.Cycle {
		index = index % len(a.Frames)
		if index < 0 { // stupid C-style modulo
			index += len(a.Frames)
		}
	} else {
		if index < 0 {
			index = 0
		}
		if index >= len(a.Frames) {
			index = len(a.Frames) - 1
		}
	}

	return a.Frames[index]
}
