package helper

import "os"

func Init_Folder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func First_Run_Init() {
	// init folder
	Init_Folder("./data")
	Init_Folder("./data/config")
	Init_Folder("./data/vms")

	// init config file
	if _, err := os.Stat("./data/config/config.json"); os.IsNotExist(err) {
		Write_Json("./data/config/config.json", map[string]interface{}{
			"vm_list": []string{},
			"vm_uuid": map[string]string{},
		})
	}
}
