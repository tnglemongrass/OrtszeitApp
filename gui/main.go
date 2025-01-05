package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
a := app.New()
	w := a.NewWindow("Ortszeit")

	hello := widget.NewLabel("Hello OrtszeitApp!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome to the Android version!")
		}),
	))

	w.ShowAndRun()
}
