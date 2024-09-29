package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Learning Time!")
	c := w.Canvas()
	w.Resize(fyne.NewSize(500, 500))

	views := ScreenViews{
		colorsView(w, c),
		numbersView(w, c),
		lettersView(w, c),
	}

	root(w, c, views)

	w.ShowAndRun()
}
