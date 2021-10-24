package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//A POC for a tabs with content view
func Mockup() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mockup POC")
	myWindow.Resize(fyne.NewSize(600, 400))

	tabs := container.NewAppTabs(
		container.NewTabItem("Addy", widget.NewLabel("Hello")),
		container.NewTabItem("Buler", widget.NewLabel("World!")),
		container.NewTabItem("Carrie", getTable()),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func getTable() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return 5, 12
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Re numbers")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			//var val int = rand.Intn(10000)
			o.(*widget.Label).SetText("number")
		})

	return table
}
