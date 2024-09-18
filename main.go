package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	oldText := "Hello Fyne!"
	newText := "Welcome :)"

	label := widget.NewLabel(oldText)
	popup := widget.ShowModalPopUp(container.NewCenter(
		widget.NewLabel("I'm a popup!"),
	))
	w.SetContent(
		container.NewVBox(
			label,
			widget.NewButton("Click Me", func() {
				if label.Text == oldText {
					label.SetText(newText)
				} else {
					label.SetText(oldText)
				}
			}),
		))

	w.ShowAndRun()
}
