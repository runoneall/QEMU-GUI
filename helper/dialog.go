package helper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ShowInfo(window fyne.Window, message string) {
	dialog.ShowCustom(
		"Info",
		"OK",
		container.NewHBox(
			widget.NewIcon(theme.InfoIcon()),
			widget.NewLabel(message),
		),
		window,
	)
}

func ShowWarning(window fyne.Window, message string) {
	dialog.ShowCustom(
		"Warning",
		"OK",
		container.NewHBox(
			widget.NewIcon(theme.WarningIcon()),
			widget.NewLabel(message),
		),
		window,
	)
}

func ShowError(window fyne.Window, message string) {
	dialog.ShowCustom(
		"Error",
		"OK",
		container.NewHBox(
			widget.NewIcon(theme.ErrorIcon()),
			widget.NewLabel(message),
		),
		window,
	)
}

func ShowFilePicker(w fyne.Window, callback func(string)) {
	dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			callback("")
			return
		}
		if f == nil {
			callback("")
			return
		}
		callback(f.URI().Path())
	}, w)
}
