package qemu_manager

import (
	"encoding/json"
	"path/filepath"
	"qemu-gui/helper"
	"qemu-gui/vars"

	"github.com/hugelgupf/vmtest/qemu"
)

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

type VMConfigExtra struct {
	Machine string `json:"Machine"`
	QEMU    string `json:"QEMU"`
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
	Extra           VMConfigExtra  `json:"Extra"`
}

func (vc *VMConfig) SaveJson() error {
	jsonData, err := json.MarshalIndent(vc, "", "  ")
	if err != nil {
		return err
	}
	path := filepath.Join(vars.CONFIG_PATH, vc.UUID+".json")
	helper.WriteFile(path, jsonData)
	return nil
}

func (vc *VMConfig) ToString() string {
	jsonData, err := json.MarshalIndent(vc, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func GetVMConfig(uuid string) (VMConfig, error) {
	path := filepath.Join(vars.CONFIG_PATH, uuid+".json")
	jsonData, err := helper.ReadFile(path)
	if err != nil {
		return VMConfig{}, err
	}
	var vc VMConfig
	err = json.Unmarshal(jsonData, &vc)
	if err != nil {
		return VMConfig{}, err
	}
	return vc, nil
}
