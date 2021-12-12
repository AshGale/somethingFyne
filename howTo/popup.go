package howTo

import (
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
	button := widget.NewButton("create popup",
		func() { go popupButton(myApp, label) })

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), button, label))
	myWindow.ShowAndRun()
}

func popupButton(myApp fyne.App, label *widget.Label) {
	log.Println("Button pressed")
	text := "before"

	//wait till the popup is finished
	waiterText := make(chan string)
	showPopUp(myApp, waiterText)
	//wait till the popup is finished
	text = <-waiterText
	label.Text = label.Text + "|" + text
	label.Refresh()
	log.Printf("Returned From popup '%+q'", text)
}

func showPopUp(a fyne.App, waiterText chan string) {
	//https://developer.fyne.io/tour/basics/windows.html
	win := a.NewWindow("Popup window example")
	textInput := widget.NewEntry()
	textInput.SetPlaceHolder("Enter text...")

	popupButton := widget.NewButton("Close", func() {
		log.Printf("sending to Main screen: %q", textInput.Text)
		waiterText <- textInput.Text
		win.Close()
	})

	win.SetContent(container.New(layout.NewVBoxLayout(), textInput, popupButton))
	win.Resize(fyne.NewSize(200, 100))
	//	win.SetOnClosed(func() { waiterText <- textInput.Text })
	win.Show()
}
