package main

import (
	"context"
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/src
var assets embed.FS

func main() {
	app := NewApp()

	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	err := wails.Run(&options.App{
		Title:  "Markdown Reader",
		Width:  1440,
		Height: 870,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx, filePath)
			if filePath != "" {
				runtime.EventsEmit(ctx, "file-opened", filePath)
			}
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}