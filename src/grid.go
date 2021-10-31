package src

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

//Example of how to do a grid
func Grid() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")
	myWindow.Resize(fyne.NewSize(600, 400))

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()

	// img := canvas.NewImageFromResource(theme.FyneLogo())
	// text := canvas.NewText("Overlay", color.Black)
	// content := container.New(layout.NewMaxLayout(), img, text)

	// myWindow.SetContent(content)
	// myWindow.ShowAndRun()
}
