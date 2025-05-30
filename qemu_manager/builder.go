package qemu_manager

import (
	"fmt"
	"math/rand"
	"net"

	"github.com/hugelgupf/vmtest/qemu"
	"github.com/hugelgupf/vmtest/qemu/qfirmware"
	"github.com/hugelgupf/vmtest/qemu/qnetwork"
)

func (vc *VMConfig) BuildOption() (*qemu.Options, error) {
	fns := []qemu.Fn{
		qemu.WithQEMUCommand(vc.WithQEMUCommand),
	}

	// name and uuid
	fns = append(fns, qemu.WithQEMUArgs(
		"-name", vc.Name,
		"-uuid", vc.UUID,
	))

	// build cpu
	fns = append(fns, qemu.WithQEMUArgs(
		"-cpu", vc.CPU.Model,
		"-smp", fmt.Sprintf(
			"cores=%s,threads=%s,sockets=1",
			vc.CPU.Cores, vc.CPU.Threads,
		),
	))

	// build memory
	fns = append(fns, qemu.WithQEMUArgs(
		"-m", fmt.Sprintf(
			"size=%s,slots=%s,maxmem=%s",
			vc.Memory.Size,
			vc.Memory.Slots,
			vc.Memory.Max,
		),
	))

	// build machine with: uefi acpi accel extra
	fns_machine := "type=" + vc.Machine
	if vc.UseUEFI {
		fns = append(fns, qfirmware.WithDefaultOVMF())
		fns_machine += ",smm=on"
	}
	if vc.UseACPI {
		fns_machine += ",acpi=on"
	}
	fns_machine += ",accel=" + vc.Accel
	if vc.Extra.Machine != "" {
		fns_machine += "," + vc.Extra.Machine
	}
	fns = append(fns, qemu.WithQEMUArgs(
		"-machine", fns_machine,
	))

	// build disk
	fns = append(fns, qemu.IDEBlockDevice(
		vc.DiskPath(),
	))
	if vc.Disk.CDROM != "" {
		fns = append(fns, qemu.WithQEMUArgs(
			"-cdrom", vc.Disk.CDROM,
		))
	}
	fns = append(fns, qemu.WithQEMUArgs(
		"-boot", "order=cd",
	))

	// build network
	fns = append(fns, qnetwork.New(
		qnetwork.WithDevice[qnetwork.UserBackend](
			qnetwork.WithNIC(qnetwork.NICVirtioNet),
			qnetwork.WithMAC(
				net.HardwareAddr{
					0x52, 0x54, 0x00,
					byte(rand.Intn(256)),
					byte(rand.Intn(256)),
					byte(rand.Intn(256)),
				},
			),
		),
	))

	// build gpu
	if vc.GPU != "virtio-gpu-pci" {
		fns = append(fns, qemu.WithQEMUArgs(
			"-vga", vc.GPU,
		))
	} else {
		fns = append(fns, qemu.WithQEMUArgs(
			"-device", vc.GPU,
		))
	}

	// build accel
	fns = append(fns, qemu.WithQEMUArgs(
		"-accel", vc.Accel,
	))

	// extra options
	fns = append(fns, qemu.WithQEMUArgs(
		// audio devices
		"-device", "ich6",
		"-device", "hda",
		"-device", "hdaudio",

		// usb devices
		"-device", "usb-ehci",
		"-device", "usb-tablet",
		"-device", "usb-mouse",
		"-device", "usb-kbd",
	))
	fns = append(fns, qemu.VirtioRandom())
	if vc.Extra.QEMU != "" {
		fns = append(fns, qemu.WithQEMUArgs(
			vc.Extra.QEMU,
		))
	}

	return qemu.OptionsFor(vc.OptionsForArch, fns...)
}
