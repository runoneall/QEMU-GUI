package qemu_manager

import "github.com/hugelgupf/vmtest/qemu"

type VMConfigCPU struct {
	Model   string `json:"Model"`
	Cores   string `json:"Cores"`
	Threads string `json:"Threads"`
}

type VMConfigMemory struct {
	Size  string `json:"Size"`
	Slots string `json:"Slots"`
	Max   string `json:"Max"`
}

type VMConfigDisk struct {
	Size  string `json:"Size"`
	CDROM string `json:"CDROM"`
}

type VMConfig struct {
	UUID            string         `json:"UUID"`
	Name            string         `json:"Name"`
	WithQEMUCommand string         `json:"WithQEMUCommand"`
	OptionsForArch  qemu.Arch      `json:"OptionsForArch"`
	CPU             VMConfigCPU    `json:"CPU"`
	Memory          VMConfigMemory `json:"Memory"`
	Machine         string         `json:"Machine"`
	UseUEFI         bool           `json:"UseUEFI"`
	UseACPI         bool           `json:"UseACPI"`
	Disk            VMConfigDisk   `json:"Disk"`
	GPU             string         `json:"GPU"`
	Accel           string         `json:"Accel"`
}
