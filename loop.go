package gogame

import "time"

// LoopFunc is a type of a game loop body function.
type LoopFunc func(ctx Context)

// Loop opens a game window based on the provided config and starts a game loop.
// At each iteration of the game loop, it calls the provided LoopFunc.
func Loop(cfg Config, lf LoopFunc) error {
	window, err := makeWindow(cfg)
	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err := makeRenderer(cfg, window)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	input := newSdlInput(window)
	output := newSdlOutput(window, renderer)

	timer := time.Now()

	var framerate <-chan time.Time
	if !cfg.VSync && cfg.FPS != 0 {
		framerate = time.Tick(time.Second / time.Duration(cfg.FPS))
	}

	for {
		input.update()

		dt := float64(time.Now().Sub(timer)) / float64(time.Second)
		timer = time.Now()

		shouldQuit := false

		lf(Context{
			Dt:       dt,
			Input:    input,
			Output:   output,
			quitFunc: func() { shouldQuit = true },
		})

		if shouldQuit {
			return nil
		}

		renderer.Present()

		if framerate != nil {
			<-framerate
		}
	}
}
