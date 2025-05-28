package helper

import (
	"os"
	"qemu-gui/vars"
)

func Init_Folder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func First_Run_Init() {
	// init folder
	Init_Folder(vars.DATA_PATH)
	Init_Folder(vars.CONFIG_PATH)
	Init_Folder(vars.VM_PATH)

	// init config file
	config_file_path := vars.CONFIG_PATH + "/config.json"
	if _, err := os.Stat(config_file_path); os.IsNotExist(err) {
		Write_Json(config_file_path, map[string]interface{}{
			"vm_list": []string{},
			"vm_uuid": map[string]string{},
		})
	}
}
