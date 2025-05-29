package qemu_manager

import (
	"fmt"
	"path/filepath"
	"qemu-gui/vars"
	"reflect"
	"strings"
)

func MapToVMConfig(m map[string]interface{}) (*VMConfig, error) {
	cfg := &VMConfig{}

	// 字符串类型字段
	strFields := []string{
		"UUID", "Name", "BootDevice", "ISO", "IMG", "Arch",
		"Machine", "Memory", "CPU", "ShareFolder", "Accel",
		"Display", "GPU", "USB", "MachineExtra", "ExtraArgs",
	}
	for _, field := range strFields {
		if val, exists := m[field]; exists {
			if cfgVal, ok := val.(string); ok {
				reflect.ValueOf(cfg).Elem().FieldByName(field).SetString(cfgVal)
			} else if val == nil {
				// 处理空值情况
				reflect.ValueOf(cfg).Elem().FieldByName(field).SetString("")
			} else {
				return nil, fmt.Errorf("field %s has invalid type, expected string", field)
			}
		}
	}

	// 布尔类型字段
	boolFields := []string{
		"UEFIBoot", "RNGDevice", "BalloonDevice", "TPM2Device",
		"ForcePS2Controller", "ClipboardShare", "FolderShare",
	}
	for _, field := range boolFields {
		if val, exists := m[field]; exists {
			if cfgVal, ok := val.(bool); ok {
				reflect.ValueOf(cfg).Elem().FieldByName(field).SetBool(cfgVal)
			} else {
				return nil, fmt.Errorf("field %s has invalid type, expected bool", field)
			}
		}
	}

	return cfg, nil
}

func BuildQEMUCommand(config *VMConfig) []string {
	cmd := []string{config.Arch}

	// 基本参数
	cmd = append(cmd, "-name", config.Name)
	cmd = append(cmd, "-uuid", config.UUID)

	// 机器类型
	machine := config.Machine
	if config.MachineExtra != "" {
		machine += "," + config.MachineExtra
	}
	cmd = append(cmd, "-machine", machine)

	// 内存和CPU
	cmd = append(cmd, "-m", config.Memory)
	cmd = append(cmd, "-smp", config.CPU)

	// 启动设备
	switch config.BootDevice {
	case "iso":
		cmd = append(cmd, "-cdrom", config.ISO)
		cmd = append(cmd, "-boot", "d")
	case "img":
		cmd = append(cmd, "-hda", config.IMG)
		cmd = append(cmd, "-boot", "c")
	}

	// 磁盘设置 - 使用VM_PATH目录下的UUID命名的qcow2文件
	diskImagePath := filepath.Join(vars.VM_PATH, config.UUID+".qcow2")
	cmd = append(cmd, "-drive", fmt.Sprintf("file=%s,format=qcow2,if=virtio", diskImagePath))

	// 加速设置
	if config.Accel != "none" {
		cmd = append(cmd, "-accel", config.Accel)
	}

	// 显示设置 - 遵循表单选择
	if config.Display != "none" {
		cmd = append(cmd, "-display", config.Display)
	}

	// GPU设置 - 特殊处理 virtio-gpu-pci
	if config.GPU != "none" {
		if config.GPU == "virtio-gpu-pci" {
			// 对于 virtio-gpu-pci，禁用标准 VGA 并使用设备
			cmd = append(cmd, "-vga", "none")
			cmd = append(cmd, "-device", "virtio-gpu-pci")
		} else {
			// 其他显卡模型使用 -vga 参数
			cmd = append(cmd, "-vga", config.GPU)
		}
	} else {
		// 无 GPU
		cmd = append(cmd, "-vga", "none")
	}

	// USB支持
	if config.USB != "none" {
		cmd = append(cmd, "-device", config.USB)
	}

	// UEFI启动
	if config.UEFIBoot {
	}

	// 设备选项
	if config.RNGDevice {
		cmd = append(cmd, "-device", "virtio-rng-pci")
	}
	if config.BalloonDevice {
		cmd = append(cmd, "-device", "virtio-balloon")
	}
	if config.TPM2Device {
		cmd = append(cmd, "-chardev", "socket,id=chrtpm,path=/tmp/tpm0/swtpm-sock",
			"-tpmdev", "emulator,id=tpm0,chardev=chrtpm",
			"-device", "tpm-tis,tpmdev=tpm0")
	}
	if config.ForcePS2Controller {
		cmd = append(cmd, "-device", "isa-ps2")
	}

	// SPICE共享功能 - 仅当启用时才添加
	if config.ClipboardShare || config.FolderShare {
		// 启用SPICE服务器
		cmd = append(cmd, "-spice", "port=5900,addr=127.0.0.1,disable-ticketing=on")

		// 基础设备
		cmd = append(cmd,
			"-device", "virtio-serial",
			"-chardev", "spicevmc,id=vdagent0,name=vdagent",
			"-device", "virtserialport,chardev=vdagent0,name=com.redhat.spice.0",
		)

		// 剪贴板共享
		if config.ClipboardShare {
			cmd = append(cmd,
				"-device", "virtio-serial",
				"-chardev", "spicevmc,id=clipboard0,name=vdagent",
				"-device", "virtserialport,chardev=clipboard0,name=com.redhat.spice.0",
			)
		}

		// 文件夹共享
		if config.FolderShare && config.ShareFolder != "" {
			cmd = append(cmd,
				"-device", "virtio-serial",
				"-chardev", "spiceport,name=org.spice-space.webdav.0,id=webdav0",
				"-device", "virtserialport,chardev=webdav0,name=org.spice-space.webdav.0",
			)
		}
	}

	// 额外参数
	if config.ExtraArgs != "" {
		args := strings.Fields(config.ExtraArgs)
		cmd = append(cmd, args...)
	}

	return cmd
}
