package main

import (
	"fynetest/constants"
	"fynetest/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Learning Time!")
	c := w.Canvas()
	w.Resize(fyne.NewSize(500, 500))

	screenViews := constants.ScreenViews{
		views.ColorsView(w, c),
		views.NumbersView(w, c),
		views.LettersView(w, c),
	}

	views.Root(w, c, screenViews)

	w.ShowAndRun()
}
