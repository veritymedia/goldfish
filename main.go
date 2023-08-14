package main

import (
	"context"
	"embed"
	"fmt"

	appicon "github.com/veritymedia/goldfish/build"
	"github.com/veritymedia/goldfish/pkg/data"
	"github.com/veritymedia/goldfish/pkg/models"
	"github.com/wailsapp/wails/v2"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/build
var assets embed.FS

type Models struct {
	ClipboardItem models.ClipboardItemModel
}

func main() {
	var err error

	// TODO: All the stuff which should be done on first app launch.
	db, err := data.InitialiseDb()
	if err != nil {
		fmt.Println("data.InitialiseDb FAILED")
	}
	// defer data.Pool.Close()

	dbModels := Models{
		ClipboardItem: models.ClipboardItemModel{DB: db},
	}

	// Create an instance of the app structure
	app := NewApp(dbModels)
	// env := NewEnv(dbModels)

	// fmt.Println("Icon bytes: ", appicon.AppIconBytes)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Goldfish",
		Width:     616,
		Height:    744,
		Frameless: true,

		// StartHidden: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// OnStartup:        app.startup,
		OnStartup: func(ctx context.Context) {
			// env.ctx = ctx
			app.ctx = ctx
			app.startClipboardWatcher()
		},
		Linux: &linux.Options{
			Icon: appicon.AppIconBytes,
		},
		OnDomReady: app.domready,
		Bind: []interface{}{
			app,
			// env,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
