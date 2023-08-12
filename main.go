package main

import (
	"embed"
	"fmt"

	"github.com/veritymedia/goldfish/pkg/data"
	"github.com/veritymedia/goldfish/pkg/models"
	"github.com/wailsapp/wails/v2"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
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

	models := Models{
		ClipboardItem: models.ClipboardItemModel{DB: db},
	}

	// str := utils.DataDir()
	// Create an instance of the app structure
	app := NewApp(models)

	/* TODO: on first launch, ask if user want app to run
	hidden in system tray. Show instructions to reset this as some
	os (like zorin) dont have a bloody system tray, making my life a pain.

	*/

	// items, err := app.Models.ClipboardItem.All()
	// if err != nil {
	// 	fmt.Println("ERROR could not get items from DB.", err)
	// }
	// fmt.Println("SUCCESS DB: got data from db: ", items)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Goldfish",
		Width:     800,
		Height:    600,
		Frameless: true,

		// StartHidden: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.domready,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
