package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}
func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := req.URL.Path
	// remove first slash
	file_path := "." + requestedFilename
	fileData, err := os.ReadFile(file_path)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", file_path)))
	}
	res.Header().Set("Content-Type", "image/svg+xml")
	res.Header().Set("Content-Disposition", "attachment; filename="+requestedFilename)
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(fileData)))
	res.Write(fileData)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "darkburn",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
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
