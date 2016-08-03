package gogame

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

// Init initializes Gogame and locks OS thread (which is necessary).
// Calling this function is necessary before using Gogame.
//
// You can call Quit on the return value, which can be combined into a nice defer:
//   defer gogame.Init().Quit()
func Init() interface {
	Quit()
} {
	runtime.LockOSThread()
	sdl.Init(sdl.INIT_EVERYTHING)
	return quitter{}
}

// Quit deinitializes Gogame and unlocks OS thread.
// Call this function when you are done with Gogame (usually on exit).
func Quit() {
	sdl.Quit()
	runtime.UnlockOSThread()
}

type quitter struct{}

func (q quitter) Quit() {
	Quit()
}
