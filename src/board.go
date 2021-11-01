package src

import (
	"fmt"
	"somethingFyne/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getTopSpacer(width, hight float32) fyne.CanvasObject {
	// size := fyne.NewSize(width, hight)
	// position := fyne.NewPos(0,0)
	// spacer := layout.Spacer{true, true, size, position ,false}
	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(width, hight))
	return spacer
}

func getBoardTiles() *fyne.Container {
	board := container.NewVBox() //use layout that relys on MinSize

	//here you figure out the days that you will show a
	//var from, till int = util.FirstDayShown, (util.FirstDayShown + util.DaysShown)
	var from, till int = util.FirstDayShown, (util.FirstDayShown + 1) //for testing just show 1 day

	for day := from; day < till; day++ {
		board.Add(getDaysTiles(day))
	}

	return board
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
			return 1, util.TimesShown
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
	dateOffset := util.FirstDayShown
	dates := make([]string, util.DaysShown)
	var text = ""

	for i := 0; i < len(dates); i++ {
		text = fmt.Sprintf(" %dth", i+dateOffset)
		dates[i] = text
	}

	table := widget.NewTable(
		func() (int, int) {
			return util.DaysShown, 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Dates")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dates[i.Row])
		})

	return table
}

func getBoard() *fyne.Container {
	//img := canvas.NewImageFromResource(theme.FyneLogo())
	board := getBoardTiles()
	//return container.New(layout.NewMaxLayout(), board)
	return container.New(layout.NewMaxLayout(), container.NewHScroll(board))
	//return board

}
