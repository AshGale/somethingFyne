package src

import (
	"somethingFyne/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//A POC for a tabs with content view
func Mockup() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mockup POC")
	myWindow.Resize(util.DefaultWindowSize)

	tabs := container.NewAppTabs(
		container.NewTabItem("Addy", getTabContent()),
		container.NewTabItem("Buler", widget.NewLabel("World!")),
		container.NewTabItem("Carrie", widget.NewLabel("Hello")),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func getTabContent() *fyne.Container {

	//replace to be inside the scroll layout
	left := getDatesTable()
	topTimes := getTimesTable(util.FirstTimeShown, util.TimesShown)
	topSpacer := getTopSpacer(left.Size().Width, topTimes.Size().Height)
	top := container.New(layout.NewBorderLayout(nil, nil, topSpacer, nil), topSpacer, topTimes)

	board := getBoard()

	return container.New(layout.NewBorderLayout(top, nil, left, nil), top, left, board)
}
