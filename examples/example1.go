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
		FPS:         60,
		QuitOnClose: true,
	}

	gogame.Loop(cfg, func(ctx gogame.Context) {
		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawLine(100, 100, 700, 500, 10, gogame.Colors["red"])
	})
}
