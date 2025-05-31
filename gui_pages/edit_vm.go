package gui_pages

import (
	"qemu-gui/helper"
	"qemu-gui/qemu_manager"
	"qemu-gui/ui_extra"
	"qemu-gui/vars"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Edit_VM_Page(myApp fyne.App, target_vm_uuid string, on_finish func()) {

	// vm window
	editVMWindow := myApp.NewWindow("New VM")
	editVMWindow.Resize(fyne.NewSize(600, 400))

	// get target vm config
	target_vm_config, err := qemu_manager.GetVMConfig(target_vm_uuid)
	if err != nil {
		helper.ShowError(editVMWindow, "Failed to get target VM config")
		go func() {
			time.Sleep(3 * time.Second)
			editVMWindow.Close()
		}()
		return
	}

	// generate vm uuid
	vm_uuid := widget.NewEntry()
	vm_uuid.SetText(target_vm_config.UUID)
	vm_uuid.Disable()

	// set vm name
	vm_name := widget.NewEntry()
	vm_name.SetText(target_vm_config.Name)

	// vm arch
	vm_arch := widget.NewSelect(vars.QEMU_SUPPORTED_ARCH, func(s string) {})
	vm_arch.SetSelected(helper.InvertMap(vars.QEMU_ARCH)[target_vm_config.WithQEMUCommand])

	// vm cpu
	vm_cpu_model := widget.NewSelectEntry(vars.QEMU_CPU)
	vm_cpu_model.SetText(target_vm_config.CPU.Model)
	vm_cpu_cores := widget.NewEntry()
	vm_cpu_cores.SetText(target_vm_config.CPU.Cores)
	vm_cpu_threads := widget.NewEntry()
	vm_cpu_threads.SetText(target_vm_config.CPU.Threads)
	vm_cpu_form_field := container.NewVBox(
		widget.NewLabel("Model"), vm_cpu_model,
		widget.NewLabel("Cores"), vm_cpu_cores,
		widget.NewLabel("Threads"), vm_cpu_threads,
	)

	// vm memory
	vm_memory_size := widget.NewEntry()
	vm_memory_size.SetText(target_vm_config.Memory.Size)
	vm_memory_slots := widget.NewEntry()
	vm_memory_slots.SetText(target_vm_config.Memory.Slots)
	vm_memory_max := widget.NewEntry()
	vm_memory_max.SetText(target_vm_config.Memory.Max)
	vm_memory_form_field := container.NewVBox(
		widget.NewLabel("Size"), vm_memory_size,
		widget.NewLabel("Slots"), vm_memory_slots,
		widget.NewLabel("Max"), vm_memory_max,
	)

	// vm machine
	vm_machine := widget.NewSelect(vars.QEMU_MACHINE, func(s string) {})
	vm_machine.SetSelected(target_vm_config.Machine)

	// vm use acpi
	vm_use_acpi := widget.NewCheck("(Need System Support)", func(b bool) {})
	vm_use_acpi.SetChecked(target_vm_config.UseACPI)

	// add disk
	vm_disk_size := widget.NewEntry()
	vm_disk_size.SetText(target_vm_config.Disk.Size)

	// add cd rom
	vm_cdrom_path := widget.NewEntry()
	vm_cdrom_path.SetText(target_vm_config.Disk.CDROM)
	vm_cdrom_path_choose := widget.NewButton("Choose", func() {
		ui_extra.FilePicker(editVMWindow, func(path string) {
			vm_cdrom_path.SetText(path)
		})
	})
	vm_cdrom_form_field := container.NewBorder(
		nil, nil, nil, vm_cdrom_path_choose, vm_cdrom_path,
	)

	// vm gpu
	vm_gpu := widget.NewSelect(vars.QEMU_GPU, func(s string) {})
	vm_gpu.SetSelected(target_vm_config.GPU)

	// vm accel
	vm_accel := widget.NewSelect(vars.QEMU_ACCEL, func(s string) {})
	vm_accel.SetSelected(target_vm_config.Accel)

	// qemu extra options
	vm_qemu_machine_extra := widget.NewEntry()
	vm_qemu_machine_extra.SetPlaceHolder("Add To -machine argument end")
	vm_qemu_machine_extra.SetText(target_vm_config.Extra.Machine)
	vm_qemu_extra := widget.NewEntry()
	vm_qemu_extra.SetPlaceHolder("Add To QEMU command end")
	vm_qemu_extra.SetText(target_vm_config.Extra.QEMU)
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
		if vm_config.Disk.Size != target_vm_config.Disk.Size {
			qemu_manager.ResizeDiskImage(vm_config.UUID, vm_config.Disk.Size)
		}

		// callback
		on_finish()

		// close window
		editVMWindow.Close()

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
	editVMWindow.SetContent(
		container.NewVScroll(
			form,
		),
	)
	editVMWindow.Show()

}
