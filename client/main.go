package main

import (
	"fmt"

	"github.com/Equanox/gotron"
)

func main() {
	// Create a new browser window instance
	window, err := gotron.New("UI/dist")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "ARBSUtility"

	// Start the browser window.
	// This will establish a golang <=> nodejs bridge using websockets,
	// to control ElectronBrowserWindow with our window object.
	done, err := window.Start()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic(err)

	}

	// Open dev tools must be used after window.Start
	 window.OpenDevTools()

	// Wait for the application to close
	<-done
}
