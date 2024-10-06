package views

import (
	"fynetest/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Root(w fyne.Window, c fyne.Canvas, view constants.ScreenViews) {

	w.SetContent(
		container.NewVBox(
			constants.RootLabel,
			widget.NewButton("Colors", func() {
				w.SetContent(view.Colors)
			}),
			widget.NewButton("Numbers", func() {
				w.SetContent(view.Numbers)
			}),
			widget.NewButton("Letters", func() {
				w.SetContent(view.Letters)
			}),
		))
}
