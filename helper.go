package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func init_folder(dir_path string) bool {
	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		os.Mkdir(dir_path, 0755)
		return true
	}
	return false
}

func write_json(file_path string, data map[string]interface{}) error {
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

func read_json(file_path string) (map[string]interface{}, error) {
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

func get_vm_list() []string {
	data, err := read_json("./data/config/config.json")
	if err != nil {
		return []string{}
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_list
}

func first_run_init() {
	// init folder
	init_folder("./data")
	init_folder("./data/config")
	init_folder("./data/vms")

	// init config file
	if _, err := os.Stat("./data/config/config.json"); os.IsNotExist(err) {
		write_json("./data/config/config.json", map[string]interface{}{
			"vm_list": []string{},
			"vm_uuid": map[string]string{},
		})
	}
}
