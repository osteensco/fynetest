package views

import (
	"fynetest/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func ColorsView(w fyne.Window, c fyne.Canvas) fyne.CanvasObject {
	var screen fyne.CanvasObject

	currColor := constants.Colors[2]
	colorBox := canvas.NewRectangle(currColor.Color)
	label := currColor.Label
	screen = container.NewStack(
		colorBox,
		label,
	)

	return screen
}
