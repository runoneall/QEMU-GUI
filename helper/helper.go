package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Init_Folder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func Write_Json(file_path string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(file_path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func Read_Json(file_path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InterfaceSliceToStringSlice(interfaceSlice []interface{}) ([]string, error) {
	stringSlice := make([]string, len(interfaceSlice))
	for i, v := range interfaceSlice {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("%d not string", i)
		}
		stringSlice[i] = str
	}
	return stringSlice, nil
}

func Get_VM_List() []string {
	data, err := Read_Json("./data/config/config.json")
	if err != nil {
		return []string{}
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_list
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

func Excutable_Command(cmd string) (bool, string) {
	// system shell
	var shell string
	var shellArg string

	if runtime.GOOS == "windows" {
		shell = "cmd"
		shellArg = "/C"
	} else {
		shell = "sh"
		shellArg = "-c"
	}

	// make command
	command := exec.Command(shell, shellArg, cmd)

	// get output
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	// execute command
	err := command.Run()

	// get result
	if err != nil {
		// execute failed
		errorMsg := stderr.String()
		if errorMsg == "" {
			errorMsg = err.Error()
		}
		return false, strings.TrimSpace(errorMsg)
	}

	// execute success
	return true, strings.TrimSpace(stdout.String())
}
