package src

import "time"

// https://stackoverflow.com/questions/38775414/golang-date-time-struct
type Day struct {
	Date        time.Time
	DisplayName string
	Duration    int
	// https://stackoverflow.com/questions/53419447/how-to-marshal-and-unmarshal-a-color-palette-to-json-in-go/53419969
	Tiles []Tile
}
