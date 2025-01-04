package main

import (
	"time"

	"OrtszeitApp/pkg/solartime"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Ortszeit")

	// Create a label to display the Ortszeit
	ortszeitLabel := widget.NewLabel("Fetching Ortszeit...")

	// Set the content of the window
	w.SetContent(container.NewVBox(
		ortszeitLabel,
	))

	// Placeholder for location fetching
	// For demonstration, we will use a fixed location (e.g., latitude: 37.7749, longitude: -122.4194)
	location := solartime.Location{
		Latitude:  37.7749,
		Longitude: -122.4194,
	}

	// Update the time periodically
	go func() {
		for {
			ortszeit := solartime.Calculate(location)
			ortszeitLabel.SetText(ortszeit)
			time.Sleep(time.Second)
		}
	}()

	// Show the window and run the app
	w.ShowAndRun()
}
