package src

import (
	"fmt"
	"log"
	"net/url"

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
	left := getDatesTable()
	topTimes := getTimesTable(7, 13)
	topSpacer := getTopSpacer(left.Size().Width, topTimes.Size().Height)
	top := container.New(layout.NewBorderLayout(nil, nil, topSpacer, nil), topSpacer, topTimes)

	board := getBoard()

	return container.New(layout.NewBorderLayout(top, nil, left, nil), top, left, board)
}

func getTopSpacer(width, hight float32) fyne.CanvasObject {
	// size := fyne.NewSize(width, hight)
	// position := fyne.NewPos(0,0)
	// spacer := layout.Spacer{true, true, size, position ,false}
	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(width, hight))
	return spacer
}

func getTile(index int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())

	url, _ := url.Parse("https://developer.fyne.io/api/v2.1/widget/hyperlink.html")
	hyperlink := widget.NewHyperlink(fmt.Sprintf("Task %v", index), url)
	tile.Add(hyperlink)

	toolbar := widget.NewToolbar(
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			log.Println("tile shouldn't move now")
		}),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {}),
		//widget.NewToolbarSpacer(),
		//widget.NewToolbarAction(theme.MenuIcon(), func() {}),

		widget.NewToolbarSeparator(),
	)

	//tile.Border

	tile.Add(toolbar)

	return tile
}

func getBoardTiles() *fyne.Container {
	var gridSize = fyne.NewSize(4, 4)
	board := container.New(layout.NewGridLayout(int(gridSize.Width)))

	for i := 0; i < int(gridSize.Width)*int(gridSize.Height); i++ {
		//tile := widget.NewCard("title", fmt.Sprintf("%v", i), canvas.NewImageFromResource(theme.FyneLogo()))
		tile := getTile(i)

		if i == 6 || i == 8 || i == 9 {
			board.Add(layout.NewSpacer())
		} else {
			board.Add(tile)
		}
	}

	return board
}

func getBoard() *fyne.Container {
	//img := canvas.NewImageFromResource(theme.FyneLogo())
	board := getBoardTiles()
	return container.New(layout.NewMaxLayout(), board)
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
