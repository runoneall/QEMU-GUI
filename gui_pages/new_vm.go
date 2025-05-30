package gui_pages

import (
	"qemu-gui/helper"
	"qemu-gui/qemu_manager"
	"qemu-gui/ui_extra"
	"qemu-gui/vars"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func New_VM_Page(myApp fyne.App, on_finish func()) {

	// vm window
	newVMWindow := myApp.NewWindow("New VM")
	newVMWindow.Resize(fyne.NewSize(600, 400))

	// generate vm uuid
	vm_uuid := widget.NewEntry()
	vm_uuid.SetText(uuid.New().String())
	vm_uuid.Disable()

	// set vm name
	vm_name := widget.NewEntry()
	vm_name.SetText("NewVM")

	// vm arch
	vm_arch := widget.NewSelect(vars.QEMU_SUPPORTED_ARCH, func(s string) {})
	vm_arch.SetSelected("amd64")

	// vm cpu
	vm_cpu_model := widget.NewSelectEntry(vars.QEMU_CPU)
	vm_cpu_model.SetText("max")
	vm_cpu_cores := widget.NewEntry()
	vm_cpu_cores.SetText("1")
	vm_cpu_threads := widget.NewEntry()
	vm_cpu_threads.SetText("2")
	vm_cpu_form_field := container.NewVBox(
		widget.NewLabel("Model"), vm_cpu_model,
		widget.NewLabel("Cores"), vm_cpu_cores,
		widget.NewLabel("Threads"), vm_cpu_threads,
	)

	// vm memory
	vm_memory_size := widget.NewEntry()
	vm_memory_size.SetText("1G")
	vm_memory_slots := widget.NewEntry()
	vm_memory_slots.SetText("1")
	vm_memory_max := widget.NewEntry()
	vm_memory_max.SetText("1G")
	vm_memory_form_field := container.NewVBox(
		widget.NewLabel("Size"), vm_memory_size,
		widget.NewLabel("Slots"), vm_memory_slots,
		widget.NewLabel("Max"), vm_memory_max,
	)

	// vm machine
	vm_machine := widget.NewSelect(vars.QEMU_MACHINE, func(s string) {})
	vm_machine.SetSelected("q35")

	// vm use uefi
	vm_use_uefi := widget.NewCheck("(Need System Support)", func(b bool) {})
	vm_use_uefi.SetChecked(true)

	// vm use acpi
	vm_use_acpi := widget.NewCheck("(Need System Support)", func(b bool) {})
	vm_use_acpi.SetChecked(true)

	// add disk
	vm_disk_size := widget.NewEntry()
	vm_disk_size.SetText("20G")

	// add cd rom
	vm_cdrom_path := widget.NewEntry()
	vm_cdrom_path_choose := widget.NewButton("Choose", func() {
		ui_extra.FilePicker(newVMWindow, func(path string) {
			vm_cdrom_path.SetText(path)
		})
	})
	vm_cdrom_form_field := container.NewBorder(
		nil, nil, nil, vm_cdrom_path_choose, vm_cdrom_path,
	)

	// vm gpu
	vm_gpu := widget.NewSelect(vars.QEMU_GPU, func(s string) {})
	vm_gpu.SetSelected("std")

	// vm accel
	vm_accel := widget.NewSelect(vars.QEMU_ACCEL, func(s string) {})
	vm_accel.SetSelected("tcg")

	// qemu extra options
	vm_qemu_machine_extra := widget.NewEntry()
	vm_qemu_machine_extra.SetPlaceHolder("Add To -machine argument end")
	vm_qemu_extra := widget.NewEntry()
	vm_qemu_extra.SetPlaceHolder("Add To QEMU command end")
	vm_qemu_extra_form_field := container.NewVBox(
		widget.NewLabel("Machine Extra Option"), vm_qemu_machine_extra,
		widget.NewLabel("QEMU Extra Option"), vm_qemu_extra,
	)

	// create vm
	vm_create := widget.NewButton("Create", func() {

		// build vm config
		vm_config := qemu_manager.VMConfig{
			UUID:            vm_uuid.Text,
			Name:            vm_name.Text,
			WithQEMUCommand: vars.QEMU_ARCH[vm_arch.Selected],
			OptionsForArch:  vars.QEMU_VMTEST_ARCH[vm_arch.Selected],
			CPU: qemu_manager.VMConfigCPU{
				Model:   vm_cpu_model.Text,
				Cores:   vm_cpu_cores.Text,
				Threads: vm_cpu_threads.Text,
			},
			Memory: qemu_manager.VMConfigMemory{
				Size:  vm_memory_size.Text,
				Slots: vm_memory_slots.Text,
				Max:   vm_memory_max.Text,
			},
			Machine: vm_machine.Selected,
			UseUEFI: vm_use_uefi.Checked,
			UseACPI: vm_use_acpi.Checked,
			Disk: qemu_manager.VMConfigDisk{
				Size:  vm_disk_size.Text,
				CDROM: vm_cdrom_path.Text,
			},
			GPU:   vm_gpu.Selected,
			Accel: vm_accel.Selected,
			Extra: qemu_manager.VMConfigExtra{
				Machine: vm_qemu_machine_extra.Text,
				QEMU:    vm_qemu_extra.Text,
			},
		}

		// save config
		vm_config.SaveJson()
		helper.AddVMToList(vm_config.Name, vm_config.UUID)
		vm_config.CreateDisk()

		// callback
		on_finish()

		// close window
		newVMWindow.Close()

	})

	// create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "UUID", Widget: vm_uuid},
			{Text: "Name", Widget: vm_name},
			{Text: "Arch", Widget: vm_arch},
			{Text: "CPU", Widget: vm_cpu_form_field},
			{Text: "Memory", Widget: vm_memory_form_field},
			{Text: "Machine", Widget: vm_machine},
			{Text: "UEFI", Widget: vm_use_uefi},
			{Text: "ACPI", Widget: vm_use_acpi},
			{Text: "Disk", Widget: vm_disk_size},
			{Text: "CD/DVD", Widget: vm_cdrom_form_field},
			{Text: "GPU", Widget: vm_gpu},
			{Text: "Accel", Widget: vm_accel},
			{Text: "Extra", Widget: vm_qemu_extra_form_field},
		},
		SubmitText: "Create",
		OnSubmit:   vm_create.OnTapped,
	}

	// show window
	newVMWindow.SetContent(
		container.NewVScroll(
			form,
		),
	)
	newVMWindow.Show()

}
