package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var oldText = "Hello AIR!!!!!!!!!!!!"
var newText = "IM A POPUP :)"

var label = widget.NewLabel(oldText)
var popupLabel = widget.NewLabel(newText)

type ScreenViews struct {
	colors  fyne.CanvasObject
	numbers fyne.CanvasObject
	letters fyne.CanvasObject
}
