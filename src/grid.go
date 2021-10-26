package src

import (
	"image/color"

	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

//Example of how to do a grid
func Grid() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")
	//myWindow.Resize(fyne.NewSize(600, 400))

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewMaxLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
