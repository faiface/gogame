package main

import "github.com/faiface/gogame"

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

	gogame.Loop(cfg, func(ctx gogame.Context) {
		ctx.SetMask(gogame.Color{1, 1, 1, 1})
		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawLine(gogame.Vec{X: 100, Y: 100}, gogame.Vec{X: 500, Y: 400}, 10, gogame.Colors["red"])
	})
}
