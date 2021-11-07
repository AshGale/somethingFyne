package src

import (
	"encoding/json"
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
	Heading  string
	Url      string
	Date     time.Time
	Duration int
	Color    color.Color
}

func getEmptyTile(duration int) *fyne.Container {
	tile := container.New(layout.NewVBoxLayout())

	// enlarger := canvas.NewRectangle(theme.BackgroundColor())
	durationLine := canvas.NewRectangle(theme.BackgroundColor())
	durationLine.SetMinSize(fyne.NewSize(util.TileSize.Width*float32(duration), 0))
	tile.Add(durationLine)

	return tile
}

func makeTileContainer(tile Tile) *fyne.Container {
	tileContainer := container.New(layout.NewVBoxLayout())

	// enlarger := canvas.NewRectangle(theme.BackgroundColor())
	durationLine := canvas.NewRectangle(util.GreenColor)
	durationLine.SetMinSize(fyne.NewSize(util.TileSize.Width*float32(tile.Duration), util.TileLineHeight))
	tileContainer.Add(durationLine)

	//url, _ := url.Parse("https://developer.fyne.io/api/v2.1/widget/hyperlink.html")
	url, _ := url.Parse(tile.Url)
	hyperlink := widget.NewHyperlink(tile.Heading, url)
	tileContainer.Add(hyperlink)

	toolbar := widget.NewToolbar(
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Printf("Edit %s\n", tile.Heading)
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			log.Printf("lock %s in time", tile.Heading)
		}),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MenuIcon(), func() {}),

		widget.NewToolbarSeparator(),
	)
	tileContainer.Add(toolbar)

	return tileContainer
}

type test struct {
	One  string
	Two  int
	Date time.Time
}

func getDaysTileData(day int) []Tile {
	//this is where you'll have to get the data for tile for the given date and time, for a person
	taskTile := "Title"
	url := "https://developer.fyne.io/api/v2.1/widget/hyperlink.html"
	tileDuration := 1
	dateData := time.Date(2021, time.December, 4, 9, 0, 0, 0, time.UTC)
	color := util.GreenColor

	tile := Tile{Heading: taskTile, Url: url, Date: dateData, Duration: tileDuration, Color: color}

	var tileList []Tile

	tileList = append(tileList, tile)
	tile.Duration = 2
	tile.Date = tile.Date.Add(time.Hour * 3)
	tileList = append(tileList, tile)

	t, err := json.Marshal(tileList)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(t))

	return tileList
}

func getDaysTiles(day int) *fyne.Container {

	daysTiles := container.NewHBox() //create new layout that is this, but has wrap capabilites
	tileList := getDaysTileData(day)
	empty := false

	//run for the number of tiles for a time period, or run for a time period and get coresponding tiles
	for i, h := 0, util.FirstTimeShown; h < util.FirstTimeShown+util.TimesShown; h++ {

		if len(tileList) == i {
			h = 100
			break
		}
		var tileData Tile = tileList[i]

		//tileData := front.Value.(Tile)
		//tileData := Tile{tile: taskTile, url: url, Date: dateData, duration: tileDuration, color: theme.BackgroundColor()}
		tileData.Heading = fmt.Sprintf("%v %v", tileData.Heading, h)
		tile := makeTileContainer(tileData)
		// tile := makeTile(fmt.Sprintf("%v %v", tileData.heading, i), tileData.url, tileData.duration)

		//log.Printf("Inedex: %v tileTime: %v ", i, tileData.Date.Hour())

		//figure out if the next tile is after than current time in loop
		if tileData.Date.Hour() > h {
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
			//tileList.Remove(tileList.Front())
			i++
		}
		fmt.Printf("index %v, hour %v\n", i, h)
	}

	// if tileList.len == i {
	// 	log.Printf("No more tiles for day %v", day)
	// } else {
	// 	log.Printf("Need to Carry over Tiles to next day(%v)", day+1)
	// }

	return daysTiles
}
