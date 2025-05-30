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
	vm_uuids, err := InterfaceSliceToStringSlice(data["vm_uuid"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_uuids
}

func GetVMName(vm_uuid string) string {
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return ""
	}
	vm_list, err := InterfaceMapToStringMap(data["vm_list"].(map[string]interface{}))
	if err != nil {
		return ""
	}
	return vm_list[vm_uuid]
}

func AddVMToList(vm_name string, vm_uuid string) bool {
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return false
	}

	// add vm_name to vm_list
	vm_list, err := InterfaceMapToStringMap(data["vm_list"].(map[string]interface{}))
	if err != nil {
		return false
	}
	vm_list[vm_uuid] = vm_name

	// add vm_uuid to vm_uuids
	vm_uuids, err := InterfaceSliceToStringSlice(data["vm_uuid"].([]interface{}))
	if err != nil {
		return false
	}
	vm_uuids = append(vm_uuids, vm_uuid)

	// update config.json
	data["vm_list"] = vm_list
	data["vm_uuid"] = vm_uuids
	return WriteJson(vars.CONFIG_FILE, data) == nil
}

func DeleteVMFromList(vm_uuid string) bool {
	data, err := ReadJson(vars.CONFIG_FILE)
	if err != nil {
		return false
	}

	// delete vm_name from vm_list
	vm_list, err := InterfaceMapToStringMap(data["vm_list"].(map[string]interface{}))
	if err != nil {
		return false
	}
	delete(vm_list, vm_uuid)

	// delete vm_uuid from vm_uuids
	vm_uuids, err := InterfaceSliceToStringSlice(data["vm_uuid"].([]interface{}))
	if err != nil {
		return false
	}
	for i, uuid := range vm_uuids {
		if uuid == vm_uuid {
			vm_uuids = append(vm_uuids[:i], vm_uuids[i+1:]...)
			break
		}
	}

	// update config.json
	data["vm_list"] = vm_list
	data["vm_uuid"] = vm_uuids
	return WriteJson(vars.CONFIG_FILE, data) == nil
}

func DeleteVMConfig(vm_uuid string) bool {
	path := filepath.Join(vars.CONFIG_PATH, vm_uuid+".json")
	return os.Remove(path) == nil
}
