package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func root(w fyne.Window, c fyne.Canvas) {

	modalContent := container.NewVBox(
		popupLabel,
		closeButton(),
	)

	modal := makePopup(c, modalContent)
	modalContent.Objects[1].(*widget.Button).OnTapped = func() { modal.Hide() }

	w.SetContent(
		container.NewVBox(
			label,
			widget.NewButton("Click Me", func() {
				modal.Show()
			}),
		))
}
