package components

import "fyne.io/fyne/v2/widget"

func closeButton() *widget.Button {

	return widget.NewButton(
		"X",
		nil,
	)
}
