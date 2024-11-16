package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/src
var assets embed.FS

func main() {
	var err error
	// motherboardSerial, err := license.GetMotherboardSerial()
	// if err != nil || motherboardSerial != "07D2212_L91D104769" {
	// 	fmt.Println("Invalid Motherboard Serial")
	// 	return
	// }
	// cpuID, err := license.GetCPUID()
	// if err != nil || cpuID != "BFEBFBFF000A0653" {
	// 	fmt.Println("invalid CPU ID")
	// 	return
	// }

	fmt.Println("License is valid")
	// 	Motherboard Serial Number: 07D2212_L91D104769
	// CPU ID: BFEBFBFF000A0653
	// Disk Serial Number: HASE23420102057
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title: "PSInventory",
		// Frameless: true, // Removes the default title bar
		// WindowStartState:  options.Maximised, //maximum width
		Width:  1366,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 127, G: 138, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
