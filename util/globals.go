package util

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var DefaultWindowSize fyne.Size = fyne.NewSize(800, 600)

var GreyColor color.RGBA = color.RGBA{R: 128, G: 128, B: 128, A: 255}
var RedColor color.RGBA = color.RGBA{R: 200, G: 0, B: 0, A: 255}
var GreenColor color.RGBA = color.RGBA{R: 0, G: 200, B: 50, A: 255}
var YellowColor color.RGBA = color.RGBA{R: 200, G: 200, B: 0, A: 255}
var BlueColor color.RGBA = color.RGBA{R: 0, G: 0, B: 200, A: 255}

var TileSize fyne.Size = fyne.NewSize(200, 100)
var TileLineHeight float32 = 10

var DaysShown int = 5
var FirstDayShown int = 4

var TimesShown int = 13
var FirstTimeShown int = 8

// type Constants struct {
// 	testColor  color.Color
// 	greenColor color.Color
// }

// var global = Constants{}

// func init() {

// 	Something = "some text"

// 	GreyColor = color.NRGBA{R: 128, G: 128, B: 128, A: 255}
// 	GreenColor = color.NRGBA{R: 0, G: 200, B: 50, A: 255}

// }
