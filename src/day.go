package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"somethingFyne/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// https://stackoverflow.com/questions/38775414/golang-date-time-struct
type Day struct {
	Date        time.Time
	DisplayName string
	Duration    int
	// https://stackoverflow.com/questions/53419447/how-to-marshal-and-unmarshal-a-color-palette-to-json-in-go/53419969
	Tiles []Tile
}

//loop through tiles, and order them to be earliest first
func sortTilesByEarliest(tiles []Tile) ([]Tile, bool) {

	var numberOfTiles int = len(tiles)
	var remainingTiles int = len(tiles)
	var sortedTiles []Tile = make([]Tile, numberOfTiles)
	var bufferTile Tile = Tile{}
	var foundIndex int = -1
	var reordered bool = false

	for sti := 0; sti < numberOfTiles; sti++ {
		for ti := 0; ti < remainingTiles; ti++ {
			if (bufferTile == Tile{}) {
				bufferTile = tiles[ti]
				foundIndex = ti
			}

			date := tiles[ti].Date
			if date.Before(bufferTile.Date) {
				fmt.Printf("tile at index %v is before %v\n", ti, bufferTile.Date)
				bufferTile = tiles[ti]
				foundIndex = ti
				reordered = true
			}

			if ti != foundIndex && date.Equal(bufferTile.Date) {
				fmt.Printf("Dates equal. ")
				tiles[ti].Date = tiles[ti].Date.Add(time.Duration(tiles[ti].Duration))
				fmt.Printf("Setting %v too %v\n", tiles[ti].Heading, tiles[ti].Date)
				reordered = true
			}
		}
		fmt.Printf("-> removing index %v and adding: %v\n", foundIndex, tiles[foundIndex].Heading)
		tiles = append(tiles[:foundIndex], tiles[foundIndex+1:]...) //end save ?
		remainingTiles = len(tiles)
		fmt.Printf("%v remainingTiles\n", remainingTiles)
		sortedTiles[sti] = bufferTile
		bufferTile = Tile{} //clear after each sort
	}

	return sortedTiles, reordered
}

func checkAndAdjustForOverlap(tiles []Tile) ([]Tile, bool) {
	overLapped := false
	//last index not needed due to no more tiles after
	for index := 0; index < len(tiles)-1; index++ {
		tileEndTime := tiles[index].Date.Add(time.Duration(tiles[index].Duration))
		if tileEndTime.After(tiles[index+1].Date) {
			//endtime is after the next tile start, and needs to be shifted
			tiles[index+1].Date = tileEndTime
			overLapped = true
		}
	}

	return tiles, overLapped
}

func getDayData(day int) Day {
	var dayData Day = Day{}

	fileName := fmt.Sprintf("./db/day%v.json", day)
	file, _ := ioutil.ReadFile(fileName)
	_ = json.Unmarshal([]byte(file), &dayData)

	dayDataSorted := false
	dayDataOverLapped := false
	//sort to not overlap or be in wrong order
	dayData.Tiles, dayDataSorted = sortTilesByEarliest(dayData.Tiles)
	dayData.Tiles, dayDataOverLapped = checkAndAdjustForOverlap(dayData.Tiles)

	//save back if file contents needed adjustment
	if dayDataSorted || dayDataOverLapped {
		log.Printf("\nData %v was updated adjusted. \nSaving ... (disabled for debug)", fileName)
		// jsonDay, _ := json.MarshalIndent(dayData, "", " ")
		// _ = ioutil.WriteFile(fileName, jsonDay, os.ModePerm)
		// fmt.Printf("Day: \n %+v\n", string(jsonDay))
	}
	return dayData
}

func getDaysTiles(day int) *fyne.Container {
	daysTiles := container.NewHBox() //create new layout that is this, but has wrap capabilites
	dayData := getDayData(day)
	tileList := dayData.Tiles
	empty := false

	//run for the number of tiles for a time period, or run for a time period and get coresponding tiles
	for i, h := 0, util.FirstTimeShown; h < util.FirstTimeShown+util.TimesShown; {

		if len(tileList) == i {
			h = 100
			break
		}
		var tileData Tile = tileList[i]
		tileContainer := makeTileContainer(tileData)

		//figure out if the next tile is after than current time in loop
		if tileData.Date.Hour() > h {
			empty = true
		} else {
			empty = false
		}

		//is there is a gap that needs to be added
		if empty {
			log.Printf("\tEmpty tile at %v", h)
			daysTiles.Add(getEmptyTile(1))
			h++
		} else {
			log.Printf("\tAdd tile at time %v", h)
			daysTiles.Add(tileContainer)
			i++
			h += tileData.Duration
		}
		fmt.Printf("index %v, hour %v\n", i, h)
	}

	return daysTiles
}
