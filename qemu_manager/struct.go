package qemu_manager

import (
	"os/exec"
	"sync"
)

type CommandManager struct {
	processes map[string]*exec.Cmd
	mu        sync.Mutex
}

type VMConfig struct {
	UUID               string
	Name               string
	BootDevice         string
	ISO                string
	IMG                string
	Arch               string
	Machine            string
	Memory             string
	CPU                string
	ShareFolder        string
	Accel              string
	Display            string
	GPU                string
	USB                string
	UEFIBoot           bool
	RNGDevice          bool
	BalloonDevice      bool
	TPM2Device         bool
	ForcePS2Controller bool
	ClipboardShare     bool
	FolderShare        bool
	MachineExtra       string
	ExtraArgs          string
}
