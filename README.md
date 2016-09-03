# Gogame

A simple and useful game library in Go. Doc: https://godoc.org/github.com/faiface/gogame

```
go get github.com/faiface/gogame
```

## Requirements

This library is using SDL2 under the hood. That means that you need to have original SDL2 libraries
installed. So far, these are required:

- SDL2
- SDL2_image
- SDL2_gfx

## Short Guide

### Basics

First, let's import Gogame.

```go
import "github.com/faiface/gogame"
```

Before you use any of Gogame, you have to initialize it. If you're writing a game, that usually
happens at the beginning of a program. When you're done with Gogame, you have to quit it.

```go
func main() {
	gogame.Init()
	defer gogame.Quit()
}
```

If you've ever written a game, you surely are familiar with the "game loop" concept. Gogame manages
this for you with `gogame.Loop` function. Let's take a loop at it's signature:

```go
func Loop(cfg Config, lf LoopFunc) error
```

Loop takes `Config` as it's first argument, and a `LoopFunc` as it's second argument. `Config` is
basically a configuration of your game window and some additional settings. `LoopFunc` stands for:

```go
type LoopFunc func(ctx Context)
```

Ok, so `LoopFunc` is a function that takes a `gogame.Context`. We'll see what it means later. Gogame
will call this function at the every iteration of the game loop. `gogame.Loop` function basically
works like this (pseudocode):

```go
func Loop(cfg Config, lf LoopFunc) error {
	// make a game window

	for {
		// update input and output

		ctx := /* make context */
		lf(ctx)
	}
}
```

Let's make a `Config` and call `gogame.Loop` to start a game! For `LoopFunc` we'll use an
anonymous function, to make it easier.

```go
package main

import "github.com/faiface/gogame"

func main() {
	gogame.Init()
	defer gogame.Quit()

	cfg := gogame.Config{
		Title:       "Hello, Gogame!" // title of the game window
		Width:       1024,            // width of the game window
		Height:      768,             // height of the game window
		FPS:         60,              // our LoopFunc will be called 60 times per second
		QuitOnClose: true,            // whenever you close a window, game loop ends
	}

	gogame.Loop(cfg, func(ctx gogame.Context) {
		// nothing here yet
	})
}
```

When you run this program, you should see an empty black window with title "Hello, Gogame!". Great!

### Context

As you could see, `LoopFunc` has one argument: a `Context`. What is it? It is a way to interact with
you game window. All the input from the mouse and keyboard, and all the drawing goes through a
`Context`.

#### Input

Gogame doesn't provide a traditional event system for input, because I don't think it's the best
way to deal with input. Instead, it provides a set of methods that allow you to check the state of
your input devices. For example, you can use `ctx.KeyDown(gogame.KeyLeft)` to check if the left
arrow is pressed on you keyboard. Here's an incomplete overview of the methods. You can find all of
them in the documentation of the `gogame.Input` interface.

```go
gogame.Loop(cfg, func(ctx gogame.Context) {
	if ctx.KeyDown(gogame.KeyLeft) {
		fmt.Println("Left arrow is being pressed down!")
	}
	if ctx.KeyJustDown(gogame.KeySpace) {
		fmt.Println("You have just pressed space!")
	}
	if ctx.KeyJustUp(gogame.KeySpace) {
		fmt.Println("You have just released space!")
	}
	fmt.Println(ctx.MousePosition())
	if ctx.MouseJustDown(gogame.MouseLeft) {
		fmt.Println("Left mouse button just pressed!")
	}
	fmt.Println(ctx.WindowSize())
})
```

#### Output

You can also use `Context` to do all kinds of graphical output (audio output is coming). Here's
again an incomplete overview, you can find all of the methods in the documentation of
`gogame.Output` interface.

```go
gogame.Loop(cfg, func(ctx gogame.Context) {
	ctx.Clear(gogame.Colors["red"])
	ctx.DrawPicture(gogame.Rect{X: 100, Y: 100, W: 50, H: 50}, myDog)
	ctx.DrawLine(gogame.Vec{X: 100, Y: 100}, gogame.Vec{X: 500, Y: 400}, 10, gogame.Colors["red"])
})
```

### Pictures

Somewhere they are called 'images', elsewhere 'textures', even elsewhere 'surfaces', we call them
*pictures*. You can load a picture from a file using `gogame.LoadPicture` function like this:

```go
myDog := gogame.LoadPicture("data/mydog.png")
```

Then you can draw a picture onto the window using the `ctx.DrawPicture` method:

```go
ctx.DrawPicture(gogame.Rect{X: 100, Y: 100, W: 50, H: 50}, myDog)
```

The picture will be stretched to fit the provided rectangle precisely.

### What else is supported?

Animations, cameras, canvases, ... look up the documentation
(https://godoc.org/github.com/faiface/gogame)!

## Notes

This library is currently heavily in development alongside my game.

It is by no means complete yet and may be changing in major ways.

## Contribution

Is very welcome! You can improve the documentation, open issues, and contribute code.
Don't hesitate and do a pull request!
