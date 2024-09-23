package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func makePopup(c fyne.Canvas, content *fyne.Container) *widget.PopUp {
	var p *widget.PopUp
	p = widget.NewModalPopUp(
		content,
		c,
	)

	return p
}

func closeButton() *widget.Button {

	return widget.NewButton(
		"X",
		nil,
	)
}
