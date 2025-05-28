package gui_pages

import (
	"fmt"
	"image/color"
	"qemu-gui/helper"
	"qemu-gui/ui_extra"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Main_Page(myApp fyne.App) *fyne.Container {

	// vm control
	vmControl := container.NewVBox(
		widget.NewLabel("Select VM To Use"),
	)
	drawVMControl := func(vm_name string, vm_info map[string]string) {
		// vm tittle
		vmControl.Add(widget.NewRichTextFromMarkdown(
			fmt.Sprintf("## %s", vm_name),
		))

		// vm info
		for k, v := range vm_info {
			if k == "start" {
				continue
			}
			vmControl.Add(widget.NewLabel(fmt.Sprintf("%s: %s", k, v)))
		}

		// vm control buttons
		vmControl.Add(container.NewHBox(
			widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {
				fmt.Println("start")
			}),
			widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
				fmt.Println("stop")
			}),
			widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
				fmt.Println("edit")
			}),
			widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
				fmt.Println("delete")
			}),
		))
	}

	// vm list
	vmList := container.NewVBox()
	vm_list_refresh := func() {
		vm_list := helper.Get_VM_List()
		vmList.RemoveAll()
		if len(vm_list) > 0 {
			for _, vm_name := range vm_list {
				vm_info := helper.GET_VM_Info(helper.GET_VM_UUID(vm_name))

				// vm item container
				vmItem := ui_extra.NewClickableContainer(
					container.NewHBox(

						// vm icon
						container.NewCenter(
							widget.NewIcon(theme.ComputerIcon()),
						),

						// vm info
						widget.NewLabel(fmt.Sprintf(
							"%s\n%s CPU, %sM Memory",
							vm_name,
							vm_info["cpu"],
							vm_info["memory"],
						)),
					),
				)

				// set on tapped
				vmItem.OnTapped = func() {
					vmControl.RemoveAll()
					drawVMControl(vm_name, vm_info)
				}

				// add to vm list
				vmList.Add(container.NewBorder(
					nil, canvas.NewRectangle(color.Gray{0x99}),
					nil, nil, vmItem,
				))

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
		vmControl,
	)
	mainContainer.SetOffset(0.25)

	// return main container
	return container.NewBorder(
		topButtons,    // top
		nil, nil, nil, // disable bottom left right
		mainContainer, // content
	)
}
