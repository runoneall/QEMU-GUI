package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"qemu-gui/vars"
)

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

func InterfaceMapToStringMap(input map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)
	for key, value := range input {
		str, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("value for key '%s' is not a string", key)
		}
		result[key] = str
	}
	return result, nil
}

func GetVMList() []string {
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return []string{}
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_list
}

func GetVMUUID(vm_name string) string {
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return ""
	}
	vm_uuids, err := InterfaceMapToStringMap(data["vm_uuid"].(map[string]interface{}))
	if err != nil {
		return ""
	}
	return vm_uuids[vm_name]
}

func AddVMToList(vm_name string, vm_uuid string) bool {

	// add vm name
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return false
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return false
	}
	vm_list = append(vm_list, vm_name)
	data["vm_list"] = vm_list

	// add vm uuid
	vm_uuids, err := InterfaceMapToStringMap(data["vm_uuid"].(map[string]interface{}))
	if err != nil {
		return false
	}
	vm_uuids[vm_name] = vm_uuid
	data["vm_uuid"] = vm_uuids

	// write config
	WriteJson(vars.CONFIG_FILE, data)
	return true
}

func DeleteVMFromList(vm_name string, vm_uuid string) bool {

	// remove vm name
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return false
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return false
	}
	for i, v := range vm_list {
		if v == vm_name {
			vm_list = append(vm_list[:i], vm_list[i+1:]...)
			break
		}
	}
	data["vm_list"] = vm_list

	// remove vm uuid
	vm_uuids, err := InterfaceMapToStringMap(data["vm_uuid"].(map[string]interface{}))
	if err != nil {
		return false
	}
	delete(vm_uuids, vm_name)
	data["vm_uuid"] = vm_uuids

	// write config
	WriteJson(vars.CONFIG_FILE, data)
	return true
}

func Delete_VM_Config(vm_uuid string) bool {
	file_path := filepath.Join(vars.CONFIG_PATH, vm_uuid+".json")
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return false
	}
	os.Remove(file_path)
	return true
}
