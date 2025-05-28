package main

import (
	"qemu-gui/gui_pages"
	"qemu-gui/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	helper.First_Run_Init()

	// init window
	myApp := app.New()
	myWindow := myApp.NewWindow("QEMU GUI")
	myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.SetContent(gui_pages.Main_Page(myApp))
	myWindow.Show()
	myApp.Run()
}
