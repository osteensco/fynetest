package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var rootLabel = widget.NewLabel("It's time to learn!")

type ScreenViews struct {
	colors  fyne.CanvasObject
	numbers fyne.CanvasObject
	letters fyne.CanvasObject
}

type ColorObj struct {
	color color.Color
	label *canvas.Text
}

var colors = []ColorObj{
	{
		color.Black,
		canvas.NewText("Black", color.White),
	},
	{
		color.White,
		canvas.NewText("White", color.Black),
	},
	{
		color.RGBA{R: 0, G: 0, B: 255, A: 100}, // blue
		canvas.NewText("Blue", color.White),
	},
}
