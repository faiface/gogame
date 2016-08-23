package main

import (
	"math"
	"time"

	"github.com/faiface/gogame"
)

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Gogame",
		Width:       1024,
		Height:      768,
		FPS:         60,
		QuitOnClose: true,
	}

	start := time.Now()

	gogame.Loop(cfg, func(ctx gogame.Context) {
		x := float64(time.Now().Sub(start)) / float64(time.Second)
		mask := gogame.Color{
			R: (math.Sin(x*math.Sqrt(2)) + 1) / 2,
			G: (math.Sin(x*math.Sqrt(3)) + 1) / 2,
			B: (math.Sin(x*math.Sqrt(5)) + 1) / 2,
			A: 1,
		}
		ctx.SetMask(mask)

		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawLine(gogame.Vec{X: 100, Y: 100}, gogame.Vec{X: 500, Y: 400}, 10, gogame.Colors["red"])
	})
}
