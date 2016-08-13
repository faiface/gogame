package gogame

// Config wraps a configuration for a game window and some minor behaviour.
type Config struct {
	Title          string
	Width          int
	Height         int
	FPS            int
	Resizable      bool
	Fullscreen     bool
	Borderless     bool
	SoftwareRender bool
	VSync          bool
	QuitOnClose    bool
}
