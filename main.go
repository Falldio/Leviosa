package main

import (
	"Leviosa/pkg/db"
	"Leviosa/pkg/log"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	log.Logger.Info("Now running Leviosa!")

	dbm := db.InitDb()
	defer dbm.Db.Close()
	// go db.FetchUpdatesForAllFeeds()
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Leviosa",
		WindowStartState: options.WindowStartState(1),
		Width:            1280,
		Height:           1024,
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
		log.Logger.Error(err.Error())
		println("Error:", err.Error())
	}

}
