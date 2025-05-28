package gui_pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func About_Page(myApp fyne.App) {

	// about window
	aboutWindow := myApp.NewWindow("About")
	aboutWindow.Resize(fyne.NewSize(400, 300))

	// right area
	aboutRight := container.NewVBox(
		widget.NewLabel("Click Left Button To Use"),
	)

	// left button
	aboutLeft := container.NewVBox(
		widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
			aboutRight.RemoveAll()
			aboutRight.Add(widget.NewLabel("关于"))
		}),
		widget.NewButtonWithIcon("", theme.ComputerIcon(), func() {
			aboutRight.RemoveAll()
			aboutRight.Add(widget.NewLabelWithStyle(
				"QEMU Excutable Check",
				fyne.TextAlignCenter,
				fyne.TextStyle{Bold: true},
			))
		}),
	)

	// show window
	aboutWindow.SetContent(container.NewHBox(
		aboutLeft,
		container.NewVScroll(aboutRight),
	))
	aboutWindow.Show()

}
