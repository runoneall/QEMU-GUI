package gui_pages

import (
	"path/filepath"
	"qemu-gui/helper"
	"qemu-gui/qemu_manager"
	"qemu-gui/ui_extra"
	"qemu-gui/vars"
	"strings"

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
	vm_uuid_str := uuid.New().String()
	vm_uuid.SetText(vm_uuid_str)
	vm_uuid.Disable()

	// vm name
	vm_name := widget.NewEntry()
	vm_name.SetText("MyVM")

	// vm boot
	vm_iso := widget.NewEntry()
	vm_iso_choose_button := widget.NewButton("Select", func() {
		ui_extra.FilePicker(newVMWindow, func(path string) {
			vm_iso.SetText(path)
		})
	})
	vm_iso_container := container.NewBorder(
		nil, nil, nil, vm_iso_choose_button, vm_iso,
	)
	vm_img := widget.NewEntry()
	vm_img_choose_button := widget.NewButton("Select", func() {
		ui_extra.FilePicker(newVMWindow, func(path string) {
			vm_img.SetText(path)
		})
	})
	vm_img_container := container.NewBorder(
		nil, nil, nil, vm_img_choose_button, vm_img,
	)
	vm_boot_devices := widget.NewSelect(vars.QEMU_BOOT_DEVICES, func(s string) {
		if s == "none" {
			vm_iso_container.Hide()
			vm_img_container.Hide()
		}
		if s == "iso" {
			vm_iso_container.Show()
			vm_img_container.Hide()
		}
		if s == "img" {
			vm_iso_container.Hide()
			vm_img_container.Show()
		}
	})
	vm_boot_devices.SetSelected("iso")
	vm_boot_container := container.NewVBox(
		vm_boot_devices,
		vm_iso_container,
		vm_img_container,
	)

	// vm hardware
	vm_arch := widget.NewSelectEntry(vars.QEMU_SYSTEMS)
	vm_machine := widget.NewSelectEntry(vars.QEMU_MACHINES)
	vm_memory := widget.NewEntry()
	vm_cpu := widget.NewEntry()
	vm_arch.SetText("qemu-system-x86_64")
	vm_machine.SetText("pc-q35-9.1")
	vm_memory.SetText("1024M")
	vm_cpu.SetText("1")

	// vm storage
	vm_storage := widget.NewEntry()
	vm_storage.SetText("20G")

	// vm share folder
	vm_share_folder := widget.NewEntry()
	vm_share_folder_choose_button := widget.NewButton("Select", func() {
		ui_extra.FolderPicker(newVMWindow, func(path string) {
			vm_share_folder.SetText(path)
		})
	})
	vm_share_folder_container := container.NewBorder(
		nil, nil, nil, vm_share_folder_choose_button, vm_share_folder,
	)

	// vm usb support
	vm_usb_support := widget.NewSelectEntry(vars.QEMU_USB_SUPPORT)
	vm_usb_support.SetText("none")

	// vm accel
	vm_accel := widget.NewSelectEntry(vars.QEMU_ACCEL)
	vm_accel.SetText("none")

	// vm display
	vm_display := widget.NewSelectEntry(vars.QEMU_DISPLAY)
	vm_display.SetText("sdl")

	// vm gpu
	vm_gpu := widget.NewSelectEntry(vars.QEMU_GPU_MODELS)

	// qemu settings
	vm_qemu_uefi_boot := widget.NewCheck("UEFI Boot", func(b bool) {})
	vm_qemu_rng_device := widget.NewCheck("RNG Device", func(b bool) {})
	vm_qemu_balloon_device := widget.NewCheck("Balloon Device", func(b bool) {})
	vm_qemu_tpm2_device := widget.NewCheck("TPM 2.0 Device", func(b bool) {})
	vm_qemu_force_ps2_controller := widget.NewCheck("Force PS/2 Controller", func(b bool) {})
	vm_qemu_clipboard_share := widget.NewCheck("Clipboard Share", func(b bool) {})
	vm_qemu_folder_share := widget.NewCheck("Folder Share", func(b bool) {})
	vm_qemu_uefi_boot.SetChecked(true)
	vm_qemu_rng_device.SetChecked(true)
	vm_qemu_clipboard_share.SetChecked(true)

	// vm qemu extra args
	vm_qemu_machine_extra := widget.NewEntry()
	vm_qemu_extra := widget.NewEntry()
	vm_qemu_machine_extra.SetPlaceHolder("Add to the end of the -machine parameter")
	vm_qemu_extra.SetPlaceHolder("Add extra arguments to qemu")

	// qemu args container
	vm_qemu_container := container.NewVBox(
		vm_qemu_uefi_boot,
		vm_qemu_rng_device,
		vm_qemu_balloon_device,
		vm_qemu_tpm2_device,
		vm_qemu_force_ps2_controller,
		vm_qemu_clipboard_share,
		vm_qemu_folder_share,
		vm_qemu_machine_extra,
		vm_qemu_extra,
	)

	// create vm
	vm_create := widget.NewButton("Create", func() {

		// get values
		config := map[string]interface{}{
			"UUID":               vm_uuid.Text,
			"Name":               vm_name.Text,
			"BootDevice":         vm_boot_devices.Selected,
			"ISO":                vm_iso.Text,
			"IMG":                vm_img.Text,
			"Arch":               vm_arch.Text,
			"Machine":            vm_machine.Text,
			"Memory":             vm_memory.Text,
			"CPU":                vm_cpu.Text,
			"ShareFolder":        vm_share_folder.Text,
			"Accel":              vm_accel.Text,
			"Display":            vm_display.Text,
			"GPU":                vm_gpu.Text,
			"USB":                vm_usb_support.Text,
			"UEFIBoot":           vm_qemu_uefi_boot.Checked,
			"RNGDevice":          vm_qemu_rng_device.Checked,
			"BalloonDevice":      vm_qemu_balloon_device.Checked,
			"TPM2Device":         vm_qemu_tpm2_device.Checked,
			"ForcePS2Controller": vm_qemu_force_ps2_controller.Checked,
			"ClipboardShare":     vm_qemu_clipboard_share.Checked,
			"FolderShare":        vm_qemu_folder_share.Checked,
			"MachineExtra":       vm_qemu_machine_extra.Text,
			"ExtraArgs":          vm_qemu_extra.Text,
		}

		// make command
		vmconfig, _ := qemu_manager.MapToVMConfig(config)
		start_command := qemu_manager.BuildQEMUCommand(vmconfig)
		config["start"] = strings.Join(start_command, " ")

		// write vm config
		helper.Write_Json(
			filepath.Join(vars.CONFIG_PATH, vm_uuid_str+".json"),
			config,
		)
		helper.Add_VM_To_List(vmconfig.Name, vmconfig.UUID)

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
			{Text: "Boot", Widget: vm_boot_container},
			{Text: "Arch", Widget: vm_arch},
			{Text: "Machine", Widget: vm_machine},
			{Text: "Memory", Widget: vm_memory},
			{Text: "CPU", Widget: vm_cpu},
			{Text: "Storage", Widget: vm_storage},
			{Text: "Share", Widget: vm_share_folder_container},
			{Text: "Accel", Widget: vm_accel},
			{Text: "Display", Widget: vm_display},
			{Text: "GPU", Widget: vm_gpu},
			{Text: "USB", Widget: vm_usb_support},
			{Text: "QEMU", Widget: vm_qemu_container},
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
