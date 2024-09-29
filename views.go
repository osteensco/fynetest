package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func root(w fyne.Window, c fyne.Canvas, view ScreenViews) {

	w.SetContent(
		container.NewVBox(
			label,
			widget.NewButton("Colors", func() {
				w.SetContent(view.colors)
			}),
			widget.NewButton("Numbers", func() {
				w.SetContent(view.numbers)
			}),
			widget.NewButton("Letters", func() {
				w.SetContent(view.letters)
			}),
		))
}

func colorsView(w fyne.Window, c fyne.Canvas) fyne.CanvasObject {
	var screen fyne.CanvasObject

	colorBox := canvas.NewRectangle(color.Black)
	label := canvas.NewText("Black", color.White)
	screen = container.NewStack(
		colorBox,
		label,
	)

	return screen
}

func numbersView(w fyne.Window, c fyne.Canvas) fyne.CanvasObject {
	var screen fyne.CanvasObject
	return screen
}

func lettersView(w fyne.Window, c fyne.Canvas) fyne.CanvasObject {
	var screen fyne.CanvasObject
	return screen
}
