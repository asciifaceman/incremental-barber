package core

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/asciifaceman/incremental-barber/pkg/data"
)

// New returns a new Game construct
func New(version string) *Game {
	return &Game{
		Version:    version,
		Resolution: fyne.NewSize(768, 768),
	}
}

// Game is our central struct that the game is built off of
type Game struct {
	Version     string
	Application fyne.App
	Resolution  fyne.Size
	Tools       *data.ToolPackage
}

// Init finishes configuration and starts up the application
func (g *Game) Init() {
	g.Application = app.New()

	// manage savegame stuff here

	// but for now jump straight to main game window
	g.MainGameWindow()

	g.Run()
}

// Run launches the fyne window manager
func (g *Game) Run() {
	g.Application.Run()
}
