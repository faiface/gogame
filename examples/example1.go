package main

import "github.com/faiface/gogame"

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Gogame",
		Width:       1024,
		Height:      768,
		Resizable:   true,
		VSync:       true,
		QuitOnClose: true,
	}

	x := 0.0

	gogame.Loop(cfg, func(ctx gogame.Context) {
		if ctx.KeyDown(gogame.KeyLeft) {
			x -= 100 * ctx.Dt
		}
		if ctx.KeyDown(gogame.KeyRight) {
			x += 100 * ctx.Dt
		}

		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawPolygon([]float64{100 + x, 600, 100}, []float64{100, 300, 500}, 0, gogame.Colors["red"])
	})
}
