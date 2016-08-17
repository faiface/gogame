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
		FPS:         60,
		QuitOnClose: true,
	}

	gogame.Loop(cfg, func(ctx gogame.Context) {
		ctx.Clear(gogame.Colors["yellow"])
		ctx.DrawLine(100, 100, 500, 400, 10, gogame.Colors["red"])
	})
}
```

## Notes

This library is currently heavily in development alongside my game.

It is by no means complete yet and may be changing in major ways.

## Contribution

Is very welcome! You can improve the documentation, open issues, and contribute code.
Don't hesitate and do a pull request!
