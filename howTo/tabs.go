package howTo

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//simple example for having tabs in fyne
func Tabs() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Addy", widget.NewLabel("Hello")),
		container.NewTabItem("Buler", widget.NewLabel("World!")),
		container.NewTabItem("Carrie", widget.NewLabel("Some other text")),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
