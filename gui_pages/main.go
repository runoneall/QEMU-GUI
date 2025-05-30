package gui_pages

import (
	"image/color"
	"qemu-gui/helper"
	"qemu-gui/qemu_manager"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var IS_VM_REFRESH = 0

func Main_Page(myApp fyne.App) *fyne.Container {

	// vm control
	vmControl := container.NewVBox(
		widget.NewLabel("Select VM To Use"),
	)
	drawVMControl := func(vm_uuid string) {

		// vm config
		vmConfig, err := qemu_manager.GetVMConfig(vm_uuid)
		if err != nil {
			vmControl.Add(widget.NewLabel("Error: " + err.Error()))
			return
		}

		// vm name title
		vmControl.Add(widget.NewRichTextFromMarkdown("## " + vmConfig.Name))

		// show vm config
		showVMConfigWidget := widget.NewMultiLineEntry()
		showVMConfigWidget.SetText(vmConfig.ToString())
		vmControl.Add(container.NewHScroll(showVMConfigWidget))

	}

	// vm list
	vmList := container.NewVBox()
	vm_list_refresh := func() {
		vm_list := helper.GetVMList()
		vmList.RemoveAll()
		if len(vm_list) > 0 {
			for _, vm_uuid := range vm_list {
				vm_name := helper.GetVMName(vm_uuid)

				// vm button
				vmList.Add(container.NewHScroll(
					widget.NewButtonWithIcon(vm_name, theme.ComputerIcon(), func() {
						vmControl.RemoveAll()
						drawVMControl(vm_uuid)
					}),
				))

			}
		} else {
			vmList.Add(widget.NewLabel("Click New to create new VM."))
		}
	}
	vm_list_refresh()
	go func() {
		for {
			if IS_VM_REFRESH == 1 {
				vm_list_refresh()
				IS_VM_REFRESH = 0
			}
		}
	}()

	// top buttons
	topButtons := container.NewVBox(
		container.NewBorder(
			nil,                                   // disable top
			canvas.NewRectangle(color.Gray{0x99}), // bottom border
			nil, nil,                              // disable left right
			container.NewHBox( // top buttons

				// new vm
				widget.NewButtonWithIcon("New", theme.ContentAddIcon(), func() {
					New_VM_Page(myApp, vm_list_refresh)
				}),

				// refresh vm list
				widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
					IS_VM_REFRESH = 1
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
		container.NewVScroll(vmControl),
	)
	mainContainer.SetOffset(0.25)

	// return main container
	return container.NewBorder(
		topButtons,    // top
		nil, nil, nil, // disable bottom left right
		mainContainer, // content
	)
}
