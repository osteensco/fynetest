package constants

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var RootLabel = widget.NewLabel("It's time to learn!")

type ScreenViews struct {
	Colors  fyne.CanvasObject
	Numbers fyne.CanvasObject
	Letters fyne.CanvasObject
}

type ColorObj struct {
	Color color.Color
	Label *canvas.Text
}

var Colors = []ColorObj{
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
