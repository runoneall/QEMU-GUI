package gui_pages

import (
	"qemu-gui/helper"
	"qemu-gui/ui_extra"
	"qemu-gui/vars"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func New_VM_Page(myApp fyne.App, on_finish func()) {

	// vm window
	newVMWindow := myApp.NewWindow("New VM")
	newVMWindow.Resize(fyne.NewSize(600, 400))

	// form fields (common set)
	vm_name := widget.NewEntry()
	vm_uuid := widget.NewEntry()
	vm_qemu := widget.NewSelectEntry(vars.QEMU_SYSTEMS)
	vm_cpu := widget.NewEntry()
	vm_memory := widget.NewEntry()

	// form fields (disk pick)
	vm_disk := widget.NewEntry()
	vm_disk_choose_button := widget.NewButton("Select", func() {
		ui_extra.FilePicker(newVMWindow, func(path string) {
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

	// create vm
	vm_create := widget.NewButton("Create", func() {
		start_cmd := ""

		// check name
		if vm_name.Text == "" {
			helper.ShowError(newVMWindow, "Please enter a name for the VM")
			return
		}
		re := regexp.MustCompile(`^[A-Za-z0-9_\-\.]+$`)
		if !re.MatchString(vm_name.Text) {
			helper.ShowError(newVMWindow, "Name can only contain letters, numbers, hyphens, underscores, and periods.")
			return
		}

		// check qemu
		if vm_qemu.Text == "" {
			helper.ShowError(newVMWindow, "Please select a QEMU system")
			return
		}

		// check cpu
		if vm_cpu.Text == "" {
			helper.ShowError(newVMWindow, "Please enter the number of CPU cores")
			return
		}

		// check memory
		if vm_memory.Text == "" {
			helper.ShowError(newVMWindow, "Please enter the amount of memory in MB")
			return
		}

		// check disk
		if vm_disk.Text == "" {
			helper.ShowError(newVMWindow, "Please select a Disk")
			return
		}

		// make command
		start_cmd += vm_qemu.Text + " "
		start_cmd += "-name " + vm_name.Text + " "
		start_cmd += "-uuid " + vm_uuid.Text + " "
		start_cmd += "-m " + vm_memory.Text + "M "
		start_cmd += "-smp " + vm_cpu.Text + " "
		start_cmd += "-hda " + vm_disk.Text + " "
		start_cmd += vm_extra_args.Text

		// save vm
		helper.Write_Json(
			vars.CONFIG_PATH+"/"+vm_uuid.Text+".json",
			map[string]interface{}{
				"uuid":   vm_uuid.Text,
				"cpu":    vm_cpu.Text,
				"memory": vm_memory.Text,
				"start":  start_cmd,
			},
		)
		helper.Add_VM_To_List(vm_name.Text, vm_uuid.Text)

		// callback
		on_finish()

		// close window
		newVMWindow.Close()

	})

	// create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: vm_name},
			{Text: "VM ID", Widget: vm_uuid},
			{Text: "QEMU", Widget: vm_qemu},
			{Text: "CPU", Widget: vm_cpu},
			{Text: "Memory", Widget: vm_memory},
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
