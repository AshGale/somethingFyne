package howTo

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//PopUp example for Fyne
func PopuP() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Popup Callback Example")
	myWindow.Resize(fyne.NewSize(300, 100))

	label := widget.NewLabel("")

	button := widget.NewButton("create popup", func() {
		log.Println("Button pressed")
		text := "before"

		//https://goinbigdata.com/golang-wait-for-all-goroutines-to-finish/
		waiter := make(chan bool)
		showPopUp(myApp, &text, waiter)
		<-waiter
		label.Text = fmt.Sprintf("from popup: '%+q'", text)
		label.Refresh()

		log.Printf("Returned From popup '%+q'", text)
	})

	//container.New(layout.NewCenterLayout(), button)
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), button, label))
	myWindow.ShowAndRun()
}

func showPopUp(a fyne.App, text *string, waiter chan bool) {
	//https://developer.fyne.io/tour/basics/windows.html
	win := a.NewWindow("Popup window example")

	textInput := widget.NewEntry()
	textInput.SetPlaceHolder("Enter text...")

	popupButton := widget.NewButton("Close", func() {
		log.Println("closing popup")
		*text = textInput.Text
		waiter <- true
		win.Close()
	})

	win.SetContent(container.New(layout.NewVBoxLayout(), textInput, popupButton))
	win.Resize(fyne.NewSize(200, 100))
	win.Show()
}
