package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fyne := app.New()

	// windows
	w_main := fyne.NewWindow("QEMU Generator")

	// main window layout
	w_main.Resize(fyne.New(150, 300))
	w_main.SetContent(widget.NewLabel("Hello, Fyne!"))

	// show window
	w_main.ShowAndRun()
}
