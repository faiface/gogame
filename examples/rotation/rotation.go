package main

import (
	"math"

	"github.com/faiface/gogame"
)

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Rotation",
		Width:       1024,
		Height:      768,
		FPS:         60,
		QuitOnClose: true,
	}

	canvas := gogame.NewCanvas(300, 300)
	canvas.Clear(gogame.Colors["blue"])

	picture := canvas.Picture().Copy()

	angle := 0.0

	gogame.Loop(cfg, func(ctx gogame.Context) {
		ctx.Clear(gogame.Colors["green"])
		ctx.DrawPicture(gogame.Rect{X: 400, Y: 200, W: 300, H: 300}, picture.Rotated(angle))

		angle += ctx.Dt * math.Pi
	})
}
