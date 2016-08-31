package gogame

// Sheet slices a single picture (spritesheet, tilesheet) into a grid of frames/tiles.
// Each frame will have a size of (frameWidth, frameHeight) and they will be produced
// line-by-line (top to bottom), frame-by-frame (left to right).
func Sheet(sheet *Picture, frameWidth, frameHeight int) []*Picture {
	var frames []*Picture

	w, h := sheet.Size()

	for x := 0; x+frameWidth <= w; x += frameWidth {
		for y := 0; y+frameHeight <= h; y += frameHeight {
			frames = append(frames, sheet.Slice(x, y, frameWidth, frameHeight))
		}
	}

	return frames
}
