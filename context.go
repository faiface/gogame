package gogame

// Context combines all input and output methods and provides some useful information too.
type Context struct {
	// Dt is the time that passed since the last call to a LoopFunc.
	Dt float64

	// Input lets you do all kinds of input.
	Input

	// Output lets you do all kinds of output.
	Output

	quitFunc func()
}

// Quit shuts the game loop down.
func (ctx *Context) Quit() {
	ctx.quitFunc()
}
