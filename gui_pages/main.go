package gui_pages

import (
	"fmt"
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
	drawVMControl := func(vm_name string, vm_info map[string]interface{}) {
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

				vm_uuid := vm_info["UUID"].(string)
				if err := qemu_manager.RunCommand(vm_uuid, vm_info["start"].(string)); err != nil {
					fmt.Printf("Failed to start VM %s: %v\n", vm_uuid, err)
				}

			}),
			widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {

				vm_uuid := vm_info["UUID"].(string)
				if err := qemu_manager.StopCommand(vm_uuid); err != nil {
					fmt.Printf("Failed to stop VM %s: %v\n", vm_uuid, err)
				}

			}),
			widget.NewButtonWithIcon("Edit", theme.DocumentCreateIcon(), func() {
				fmt.Println("edit")
			}),
			widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {

				// delete vm
				vm_uuid := vm_info["UUID"].(string)
				helper.Delete_VM_From_List(vm_name, vm_uuid)
				helper.Delete_VM_Config(vm_uuid)

				// refresh vm list
				IS_VM_REFRESH = 1

				// clear vm control
				vmControl.RemoveAll()
				vmControl.Add(widget.NewLabel("Select VM To Use"))

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

				vmList.Add(container.NewHScroll(
					widget.NewButtonWithIcon(
						vm_name, theme.ComputerIcon(), func() {
							vmControl.RemoveAll()
							drawVMControl(vm_name, vm_info)
						},
					),
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
