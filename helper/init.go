package helper

import (
	"os"
	"qemu-gui/vars"
)

func InitFolder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
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
	if _, err := os.Stat(config_file_path); os.IsNotExist(err) {
		WriteJson(config_file_path, map[string]interface{}{
			"vm_list": []string{},
			"vm_uuid": map[string]string{},
		})
	}
}
