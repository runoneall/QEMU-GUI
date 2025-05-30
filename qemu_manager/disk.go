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
