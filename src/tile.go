package src

import (
	"container/list"
	"fmt"
	"image/color"
	"log"
	"net/url"
	"somethingFyne/util"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// https://stackoverflow.com/questions/38775414/golang-date-time-struct
type Tile struct {
	tile     string
	url      string
	Date     time.Time
	duration int
	color    color.Color
}

func getEmptyTile(duration int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())

	// enlarger := canvas.NewRectangle(theme.BackgroundColor())
	durationLine := canvas.NewRectangle(theme.BackgroundColor())
	durationLine.SetMinSize(fyne.NewSize(util.TileSize.Width*float32(duration), 0))
	tile.Add(durationLine)

	return tile
}

func getTile(taskTile, urlString string, duration int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())

	// enlarger := canvas.NewRectangle(theme.BackgroundColor()
	durationLine := canvas.NewRectangle(util.GreenColor)
	durationLine.SetMinSize(fyne.NewSize(util.TileSize.Width*float32(duration), util.TileLineHeight))
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

func getDaysTileData(day int) *list.List {
	//this is where you'll have to get the data for tile for the given date and time, for a perso
	taskTile := "Title"
	url := "https://developer.fyne.io/api/v2.1/widget/hyperlink.html"
	tileDuration := 1

	dateData := time.Date(2021, time.December, 4, 9, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", dateData.Local())

	tileData := Tile{tile: taskTile, url: url, Date: dateData, duration: tileDuration, color: theme.BackgroundColor()}

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

	return list
}

func getDaysTiles(day int) *fyne.Container {

	daysTiles := container.NewHBox() //create new layout that is this, but has wrap capabilites
	tileList := getDaysTileData(day)
	empty := false

	//run for the number of tiles for a time period, or run for a time period and get coresponding tiles
	for i := util.FirstTimeShown; i < util.FirstTimeShown+util.TimesShown; i++ {

		front := tileList.Front()
		if front == nil {
			i = 100
			break
		}

		tileData := front.Value.(Tile)
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
			//daysTiles.Add(layout.NewSpacer())//only works with set grid sizes
			daysTiles.Add(getEmptyTile(1))
		} else {
			log.Printf("\tAdded new tile at time %v", tileData.Date.Hour())
			//todo, might be the place to update the hour of the tile,
			daysTiles.Add(tile)
			tileList.Remove(tileList.Front())
		}
	}

	if tileList.Front() == nil {
		log.Printf("No more tiles for day %v", day)
	} else {
		log.Printf("Need to Carry over Tiles to next day(%v)", day+1)
	}

	return daysTiles
}



