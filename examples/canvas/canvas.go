package main

import "github.com/faiface/gogame"

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Canvas Example",
		Width:       1024,
		Height:      768,
		FPS:         60,
		QuitOnClose: true,
	}

	canvas := gogame.NewCanvas(cfg.Width/2, cfg.Height/2)
	canvas.Clear(gogame.Colors["white"])
	canvas.DrawPolygon([]gogame.Vec{{100, 100}, {100, 300}, {450, 200}}, 0, gogame.Colors["grey"])

	gogame.Loop(cfg, func(ctx gogame.Context) {
		outputRect := ctx.OutputRect()
		middle := outputRect.Size().D(2)

		ctx.SetMask(gogame.Colors["red"])
		ctx.DrawPicture(gogame.Rect{X: 0, Y: 0, W: middle.X, H: middle.Y}, canvas.Picture())

		ctx.SetMask(gogame.Colors["green"])
		ctx.DrawPicture(gogame.Rect{X: middle.X, Y: 0, W: middle.X, H: middle.Y}, canvas.Picture())

		ctx.SetMask(gogame.Colors["yellow"])
		ctx.DrawPicture(gogame.Rect{X: 0, Y: middle.Y, W: middle.X, H: middle.Y}, canvas.Picture())

		ctx.SetMask(gogame.Colors["blue"])
		ctx.DrawPicture(gogame.Rect{X: middle.X, Y: middle.Y, W: middle.X, H: middle.Y}, canvas.Picture())
	})
}
