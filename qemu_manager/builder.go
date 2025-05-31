package qemu_manager

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
)

func appendArgs(cmd string, args [][]string) string {
	for _, arg := range args {
		cmd += " " + strings.Join(arg, " ")
	}
	return cmd
}

func (vc *VMConfig) BuildOption() string {
	start_command := vc.WithQEMUCommand

	// name and uuid
	start_command = appendArgs(start_command, [][]string{
		{"-name", vc.Name},
		{"-uuid", vc.UUID},
	})

	// build cpu
	start_command = appendArgs(start_command, [][]string{
		{"-cpu", vc.CPU.Model},
		{"-smp", fmt.Sprintf(
			"cores=%s,threads=%s,sockets=1",
			vc.CPU.Cores, vc.CPU.Threads,
		)},
	})

	// build memory
	fns_memory := "size=" + vc.Memory.Size
	if vc.Memory.Max != vc.Memory.Size {
		fns_memory += ",slots=" + vc.Memory.Slots
		fns_memory += ",maxmem=" + vc.Memory.Max
	}
	start_command = appendArgs(start_command, [][]string{
		{"-m", fns_memory},
	})

	// build machine with: uefi acpi extra
	fns_machine := "type=" + vc.Machine
	if vc.UseACPI {
		fns_machine += ",acpi=on"
	}
	if vc.Extra.Machine != "" {
		fns_machine += "," + vc.Extra.Machine
	}
	start_command = appendArgs(start_command, [][]string{
		{"-machine", fns_machine},
	})

	// build disk
	if vc.Disk.CDROM != "" {
		start_command = appendArgs(start_command, [][]string{
			{"-cdrom", vc.Disk.CDROM},
		})
	}
	start_command = appendArgs(start_command, [][]string{
		{"-hda", vc.DiskPath()},
		{"-boot", "order=cd"},
	})

	// build network
	start_command = appendArgs(start_command, [][]string{
		{"-nic", fmt.Sprintf("user,model=virtio-net-pci,mac=%s",
			net.HardwareAddr{
				0x52, 0x54, 0x00,
				byte(rand.Intn(256)),
				byte(rand.Intn(256)),
				byte(rand.Intn(256)),
			},
		)},
	})

	// build gpu
	if vc.GPU != "virtio-gpu-pci" {
		start_command = appendArgs(start_command, [][]string{
			{"-vga", vc.GPU},
		})
	} else {
		start_command = appendArgs(start_command, [][]string{
			{"-device", vc.GPU},
		})
	}

	// build accel
	start_command = appendArgs(start_command, [][]string{
		{"-accel", vc.Accel},
	})

	// extra options
	start_command = appendArgs(start_command, [][]string{
		// usb devices
		{"-device", "usb-ehci"},
		{"-device", "usb-tablet"},
		{"-device", "usb-mouse"},
		{"-device", "usb-kbd"},

		// random devices
		{"-device", "virtio-rng-pci"},
	})
	start_command = appendArgs(start_command, [][]string{
		{"", vc.Extra.QEMU},
	})

	return start_command
}
