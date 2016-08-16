package main

import (
	"fmt"

	"github.com/faiface/gogame"
)

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:     "Gogame",
		Width:     1024,
		Height:    768,
		Resizable: true,
		FPS:       60,
	}

	gogame.Loop(cfg, func(ctx gogame.Context) {
		if ctx.WindowClosed() {
			ctx.Quit()
		}
		if ctx.KeyJustDown(gogame.KeySpace) {
			fmt.Println("space pressed")
		}
		if ctx.WindowResized() {
			fmt.Println(ctx.WindowSize())
		}
		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawLine(100, 100, 700, 500, 10, gogame.Colors["red"])
	})
}
