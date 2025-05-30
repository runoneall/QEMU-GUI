package main

import (
	"qemu-gui/gui_pages"
	"qemu-gui/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	helper.FirstRunInit()

	// init window
	myApp := app.New()
	myWindow := myApp.NewWindow("QEMU GUI")
	myWindow.Resize(fyne.NewSize(800, 600))

	// set content
	myWindow.SetContent(gui_pages.Main_Page(myApp))
	myWindow.Show()
	myApp.Run()
}
