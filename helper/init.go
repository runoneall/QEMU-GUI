package helper

import (
	"os"
	"qemu-gui/vars"
)

func InitFolder(dir_path string) bool {
	if !IsExist(dir_path) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func FirstRunInit() {
	// init folder
	InitFolder(vars.DATA_PATH)
	InitFolder(vars.CONFIG_PATH)
	InitFolder(vars.VM_PATH)

	// init config file
	config_file_path := vars.CONFIG_FILE
	if !IsExist(config_file_path) {
		WriteJson(config_file_path, map[string]interface{}{
			"vm_list": map[string]string{},
			"vm_uuid": []string{},
		})
	}
}
