package gui_pages

import (
	"fmt"
	"image/color"
	"qemu-gui/helper"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Main_Page(myApp fyne.App) *fyne.Container {

	// vm list
	vmList := container.NewVBox()
	vm_list_refresh := func() {
		vm_list := helper.Get_VM_List()
		vmList.RemoveAll()
		if len(vm_list) > 0 {
			for _, vm_name := range vm_list {
				vmList.Add(widget.NewLabel(vm_name))
			}
		} else {
			vmList.Add(widget.NewLabel("Click New to create new VM."))
		}
	}
	vm_list_refresh()

	// top buttons
	topButtons := container.NewVBox(
		container.NewBorder(
			nil,                                   // disable top
			canvas.NewRectangle(color.Gray{0x99}), // bottom border
			nil, nil,                              // disable left right
			container.NewHBox( // top buttons

				// new vm
				widget.NewButtonWithIcon("New", theme.DocumentCreateIcon(), func() {
					New_VM_Page(myApp)
				}),

				// refresh vm list
				widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
					vm_list_refresh()
				}),

				// settings
				widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
					fmt.Println("settings")
				}),

				// about
				widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
					About_Page(myApp)
				}),

				// exit
				widget.NewButtonWithIcon("Exit", theme.CancelIcon(), func() {
					myApp.Quit()
				}),

				layout.NewSpacer(),
			),
		),
	)

	// show window
	mainContainer := container.NewHSplit(
		container.NewVScroll(vmList),
		widget.NewLabel("VM Control"),
	)
	mainContainer.SetOffset(0.25)

	// return main container
	return container.NewBorder(
		topButtons,    // top
		nil, nil, nil, // disable bottom left right
		mainContainer, // content
	)
}
