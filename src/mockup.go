package src

import (
	"container/list"
	"fmt"
	"image/color"
	"log"
	"net/url"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var testColor = color.NRGBA{R: 128, G: 128, B: 128, A: 255}

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

func getTopSpacer(width, hight float32) fyne.CanvasObject {
	// size := fyne.NewSize(width, hight)
	// position := fyne.NewPos(0,0)
	// spacer := layout.Spacer{true, true, size, position ,false}
	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(width, hight))
	return spacer
}

func getEmptyTile(duration int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())
	size := fyne.NewSize(200, 100)

	// enlarger := canvas.NewRectangle(theme.BackgroundColor())
	durationLine := canvas.NewRectangle(testColor)
	durationLine.SetMinSize(fyne.NewSize(size.Width*float32(duration), 0))
	tile.Add(durationLine)

	return tile
}

func getTile(taskTile, urlString string, duration int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())
	size := fyne.NewSize(200, 100)

	// enlarger := canvas.NewRectangle(theme.BackgroundColor())
	durationLine := canvas.NewRectangle(testColor)
	durationLine.SetMinSize(fyne.NewSize(size.Width*float32(duration), 1))
	tile.Add(durationLine)

	//url, _ := url.Parse("https://developer.fyne.io/api/v2.1/widget/hyperlink.html")
	url, _ := url.Parse(urlString)
	hyperlink := widget.NewHyperlink(taskTile, url)
	tile.Add(hyperlink)

	toolbar := widget.NewToolbar(
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Printf("Edit %s\n", taskTile)
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			log.Printf("lock %s in time", taskTile)
		}),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MenuIcon(), func() {}),

		widget.NewToolbarSeparator(),
	)
	tile.Add(toolbar)

	return tile
}

// https://stackoverflow.com/questions/38775414/golang-date-time-struct
type Tile struct {
	tile     string
	url      string
	Date     time.Time
	duration int
}

func getDaysTiles() *fyne.Container {

	day := container.NewHBox() //create new layout that is this, but has wrap capabilites

	//this is where you'll have to get the data for tile for the given date and time, for a person
	taskTile := "Title"
	url := "https://developer.fyne.io/api/v2.1/widget/hyperlink.html"
	tileDuration := 1
	empty := false

	dateData := time.Date(2021, time.December, 4, 9, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", dateData.Local())

	tileData := Tile{tile: taskTile, url: url, Date: dateData, duration: tileDuration}

	list := list.New()
	list.PushBack(tileData)
	tileData.Date = tileData.Date.Add(time.Hour)
	list.PushBack(tileData)
	tileData.Date = tileData.Date.Add(time.Hour)
	list.PushBack(tileData)
	tileData.Date = tileData.Date.Add(time.Hour * 2)
	list.PushBack(tileData)
	tileData.Date = tileData.Date.Add(time.Hour)
	tileData.duration = 2
	list.PushBack(tileData)
	tileData.Date = tileData.Date.Add(time.Hour)
	tileData.duration = 1
	list.PushBack(tileData)

	//run for the number of tiles for a time period, or run for a time period and get coresponding tiles
	for i := 7; i < 20; i++ {

		front := list.Front()
		if front == nil {
			i = 100
			break
		}

		tileData = front.Value.(Tile)
		tile := getTile(fmt.Sprintf("%v %v", tileData.tile, i), tileData.url, tileData.duration)

		log.Printf("Inedex: %v tileTime: %v ", i, tileData.Date.Hour())

		//figure out if the next tile is after than current time in loop
		if tileData.Date.Hour() > i {
			empty = true
		} else {
			empty = false
		}

		//need to figure out if there is a gap that needs to be added
		if empty {
			log.Printf("\tAdded empty tile")
			//day.Add(layout.NewSpacer())//only works with set grid sizes
			day.Add(getEmptyTile(1))
		} else {
			log.Printf("\tAdded new tile at time %v", tileData.Date.Hour())
			//todo, might be the place to update the hour of the tile,
			day.Add(tile)
			list.Remove(list.Front())
		}
	}

	return day
}

func getBoardTiles() *fyne.Container {
	//var gridSize = fyne.NewSize(4, 4)
	//board := container.New(layout.NewGridLayout(int(gridSize.Width)))
	//board := container.New(layout.NewGridWrapLayout(fyne.NewSize(200, 100)))
	board := container.NewVBox() //create new layout that is this, but has wrap capabilites

	board.Add(getDaysTiles())

	return board
}

func getBoard() *fyne.Container {
	//img := canvas.NewImageFromResource(theme.FyneLogo())
	board := getBoardTiles()
	//return container.New(layout.NewMaxLayout(), board)
	return container.New(layout.NewMaxLayout(), container.NewHScroll(board))
	//return board

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

func getTabContent() *fyne.Container {

	//replace to be inside the scroll layout
	left := getDatesTable()
	topTimes := getTimesTable(7, 13)
	topSpacer := getTopSpacer(left.Size().Width, topTimes.Size().Height)
	top := container.New(layout.NewBorderLayout(nil, nil, topSpacer, nil), topSpacer, topTimes)

	board := getBoard()

	return container.New(layout.NewBorderLayout(top, nil, left, nil), top, left, board)
}
