package gui_pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func New_VM_Page(myApp fyne.App) {

	// new vm window
	newVMWindow := myApp.NewWindow("New VM")
	newVMWindow.Resize(fyne.NewSize(400, 200))

	// show window
	newVMWindow.SetContent(widget.NewLabel("New VM"))
	newVMWindow.Show()

}
