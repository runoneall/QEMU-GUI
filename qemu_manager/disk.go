package qemu_manager

import (
	"os"
	"os/exec"
	"path/filepath"
	"qemu-gui/vars"
)

func CreateDiskImage(uuid string, format string, size string) (string, error) {
	path := filepath.Join(vars.VM_PATH, uuid+"."+format)
	cmd := exec.Command("qemu-img", "create", "-f", format, path, size)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return path, cmd.Run()
}

func RemoveDiskImage(uuid string) error {
	path := filepath.Join(vars.VM_PATH, uuid+".qcow2")
	return os.Remove(path)
}

func ResizeDiskImage(uuid string, size string) error {
	path := filepath.Join(vars.VM_PATH, uuid+".qcow2")
	cmd := exec.Command("qemu-img", "resize", path, size)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (vc *VMConfig) CreateDisk() error {
	_, err := CreateDiskImage(vc.UUID, "qcow2", vc.Disk.Size)
	return err
}

func (vc *VMConfig) DiskPath() string {
	return filepath.Join(vars.VM_PATH, vc.UUID+".qcow2")
}

func (vc *VMConfig) RemoveDisk() error {
	return RemoveDiskImage(vc.UUID)
}

func (vc *VMConfig) ResizeDisk(size string) error {
	return ResizeDiskImage(vc.UUID, size)
}
