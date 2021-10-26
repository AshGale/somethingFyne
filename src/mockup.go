package src

import (
	"fmt"

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
		container.NewTabItem("Addy", getTable()),
		container.NewTabItem("Buler", widget.NewLabel("World!")),
		container.NewTabItem("Carrie", widget.NewLabel("Hello")),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func getTitleFormat(dates [][]string) [][]string {
	//for testing name all dates shown after 3rd
	dateOffset := 4
	hourOffset := 7

	text := ""

	for i := 0; i < len(dates); i++ {
		for j := 0; j < len(dates[i]); j++ {
			if j+hourOffset < 12 {
				text = fmt.Sprintf(" - %dth - \n %dam", i+dateOffset, j+hourOffset)
			} else {
				text = fmt.Sprintf(" - %dth - \n %dpm", i+dateOffset, j+hourOffset)
			}
			dates[i][j] = text
		}
	}

	return dates
}

func getTable() *widget.Table {

	//5 12

	dates := make([][]string, 5)
	for i := range dates {
		dates[i] = make([]string, 12)
	}

	dates = getTitleFormat(dates)

	table := widget.NewTable(
		func() (int, int) {
			return 5, 12
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Re numbers")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dates[i.Row][i.Col])

			//o.(*widget.Card).SetTitle(dates[i.Row][i.Col])

		})

	return table
}
