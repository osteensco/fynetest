package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var modalContent = container.NewVBox(
	popupLabel,
	closeButton(),
)

// var modal = makePopup(c, modalContent)
// modalContent.Objects[1].(*widget.Button).OnTapped = func() { modal.Hide() }

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
