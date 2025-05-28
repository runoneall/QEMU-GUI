package gui_pages

import (
	"qemu-gui/helper"
	"qemu-gui/vars"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func New_VM_Page(myApp fyne.App) {

	// vm window
	newVMWindow := myApp.NewWindow("New VM")
	newVMWindow.Resize(fyne.NewSize(600, 400))

	// form fields (common set)
	vm_name := widget.NewEntry()
	vm_uuid := widget.NewEntry()
	vm_qemu := widget.NewSelectEntry(vars.QEMU_SYSTEMS)
	vm_cpu := widget.NewEntry()
	vm_memory := widget.NewEntry()
	vm_enable_kvm := widget.NewCheck("(Need System Support)", func(checked bool) {})

	// form fields (disk pick)
	vm_disk := widget.NewEntry()
	vm_disk_choose_button := widget.NewButton("Select", func() {
		helper.ShowFilePicker(newVMWindow, func(path string) {
			if path == "" {
				vm_disk.SetText("No File Selected")
				return
			}
			vm_disk.SetText(path)
		})
	})

	// form fields (extra args)
	vm_extra_args := widget.NewEntry()

	// generate vm uuid
	vm_uuid_str := uuid.New().String()
	vm_uuid.SetText(vm_uuid_str)
	vm_uuid.Disable()

	// default qemu
	vm_qemu.SetText("qemu-system-x86_64")

	// vm cpu cores
	vm_cpu.PlaceHolder = "Cores"
	vm_cpu.SetText("1")

	// vm memory placeholder
	vm_memory.PlaceHolder = "MB"
	vm_memory.SetText("512")

	// default enable kvm
	vm_enable_kvm.SetChecked(true)

	// create vm
	vm_create := widget.NewButton("Create", func() {
		// start_cmd := ""

		// check name
		if vm_name.Text == "" {
			helper.ShowError(newVMWindow, "Please enter a name for the VM")
		}

		// check qemu
		if vm_qemu.Text == "" {
			helper.ShowError(newVMWindow, "Please select a QEMU system")
		}

		// check cpu
		if vm_cpu.Text == "" {
			helper.ShowError(newVMWindow, "Please enter the number of CPU cores")
		}

		// check memory
		if vm_memory.Text == "" {
			helper.ShowError(newVMWindow, "Please enter the amount of memory in MB")
		}

		// check disk
		if vm_disk.Text == "" {
			helper.ShowError(newVMWindow, "Please select a Disk")
		}

	})

	// create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: vm_name},
			{Text: "VM ID", Widget: vm_uuid},
			{Text: "QEMU", Widget: vm_qemu},
			{Text: "CPU", Widget: vm_cpu},
			{Text: "Memory", Widget: vm_memory},
			{Text: "Enable KVM", Widget: vm_enable_kvm},
			{Text: "Disk", Widget: container.NewBorder(
				nil, nil, nil,
				vm_disk_choose_button, vm_disk,
			)},
			{Text: "Extra Args", Widget: vm_extra_args},
		},
		SubmitText: "Create",
		OnSubmit:   vm_create.OnTapped,
	}

	// show window
	newVMWindow.SetContent(form)
	newVMWindow.Show()

}
