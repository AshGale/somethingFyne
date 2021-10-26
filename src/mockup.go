package src

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//A POC for a tabs with content view
func Mockup() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mockup POC")
	myWindow.Resize(fyne.NewSize(800, 600))

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

	//maybe replace with static text to prevent scolling
	topSpacer := getTopSpacer()
	top := container.New(layout.NewBorderLayout(nil, nil, topSpacer, nil), topSpacer, getTimesTable(7, 13))
	left := getDatesTable()
	board := getBoard()

	return container.New(layout.NewBorderLayout(top, nil, left, nil), top, left, board)
}

func getTopSpacer() *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return 1, 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("            ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("            ")
		})

	return table
}

func getBoard() *fyne.Container {
	img := canvas.NewImageFromResource(theme.FyneLogo())
	return container.New(layout.NewMaxLayout(), img)
}

func getTimesTable(hourOffset, duration int) *widget.Table {

	times := make([]string, duration)
	var text = ""

	for i := 0; i < len(times); i++ {

		if i+hourOffset < 12 {
			text = fmt.Sprintf("%dam", i+hourOffset)
		} else {
			text = fmt.Sprintf(" %dpm", i+hourOffset)
		}
		times[i] = text
	}

	table := widget.NewTable(
		func() (int, int) {
			return 1, 13
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Hours")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(times[i.Col])
		})

	return table
}

func getDatesTable() *widget.Table {
	dateOffset := 4
	dates := make([]string, 6)
	var text = ""

	for i := 0; i < len(dates); i++ {
		text = fmt.Sprintf(" %dth", i+dateOffset)
		dates[i] = text
	}

	table := widget.NewTable(
		func() (int, int) {
			return 5, 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Dates")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dates[i.Row])
		})

	return table
}
