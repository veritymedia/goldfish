package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	/* TODO: on first launch, ask if user want app to run
	hidden in system tray. Show instructions to reset this as some
	os (like zorin) dont have a bloody system tray, making my life a pain.

	*/

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "Goldfish",
		Width:       800,
		Height:      600,
		Frameless:   true,
		StartHidden: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
