# Gogame

A simple and useful game library in Go.

```go
package main

import "github.com/faiface/gogame"

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Gogame",
		Width:       1024,
		Height:      768,
		QuitOnClose: true,
	}

	gogame.Loop(cfg, func(ctx gogame.Context) {
		ctx.Clear(gogame.Yellow)
		ctx.DrawRect(100, 100, 500, 400, 0, gogame.Red)
	})
}

```
