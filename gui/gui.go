package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunGUI() {
	
	app := app.New()
	window := app.NewWindow("GolyPoly")
	window.SetContent(container.NewVBox(
		widget.NewLabel("Yoo"),
		widget.NewSlider(20, 20000),
		))

	window.ShowAndRun()
}
