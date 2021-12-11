package howTo

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

//PopUp example for Fyne
func PopuP() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Widget")
	myWindow.Resize(fyne.NewSize(150, 150))

	button := widget.NewButton("create popup", func() {
		log.Println("Button pressed")
		showPopUp(myApp)
		log.Println("Button click funtion end")
	})

	myWindow.SetContent(button)
	myWindow.ShowAndRun()
}

func showPopUp(a fyne.App) {
	//https://developer.fyne.io/tour/basics/windows.html
	win := a.NewWindow("Popup window example")

	content := widget.NewButton("close popup", func() {
		log.Println("closing popup")
		win.Close()
	})

	win.SetContent(content)
	win.Resize(fyne.NewSize(150, 150))
	win.Show()
}
