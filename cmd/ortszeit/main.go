package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"OrtszeitApp/pkg/solartime"
)

type ortszeit struct {
	window        fyne.Window
	latitudeBind  binding.Float
	longitudeBind binding.Float
	solarTimeBind binding.String
	offsetBind    binding.String
	locationLabel *widget.Label
}

func (o *ortszeit) updateLocation(lat, long float64) {
	o.latitudeBind.Set(lat)
	o.longitudeBind.Set(long)
	o.locationLabel.SetText(fmt.Sprintf("Location: %.4f, %.4f", lat, long))
}

func main() {
	myApp := app.NewWithID("com.ortszeitapp.ortszeit")
	window := myApp.NewWindow("Ortszeit")

	// Create app state
	app := &ortszeit{
		window:        window,
		latitudeBind:  binding.NewFloat(),
		longitudeBind: binding.NewFloat(),
		solarTimeBind: binding.NewString(),
		offsetBind:    binding.NewString(),
		locationLabel: widget.NewLabel("Waiting for GPS..."),
	}

	// Create display labels with larger text
	solarTimeLabel := widget.NewLabelWithData(app.solarTimeBind)
	solarTimeLabel.TextStyle = fyne.TextStyle{Bold: true}
	
	offsetLabel := widget.NewLabelWithData(app.offsetBind)

	// Update function
	updateTime := func() {
		lat, _ := app.latitudeBind.Get()
		long, _ := app.longitudeBind.Get()

		loc := solartime.Location{
			Latitude:  lat,
			Longitude: long,
		}

		now := time.Now()
		solarTime := solartime.GetSolarTime(loc, now)
		offset := solarTime.Sub(now)

		app.solarTimeBind.Set(fmt.Sprintf("Solar Time: %s", solarTime.Format("15:04:05")))
		app.offsetBind.Set(fmt.Sprintf("Offset: %.1f minutes", offset.Minutes()))
	}

	// Set default values (Munich)
	app.updateLocation(48.1351, 11.5820)

	// Initial update
	updateTime()

	// Start timer for updates
	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			updateTime()
		}
	}()

	// Layout
	content := container.NewVBox(
		solarTimeLabel,
		offsetLabel,
		widget.NewSeparator(),
		app.locationLabel,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 200))
	window.ShowAndRun()
}
