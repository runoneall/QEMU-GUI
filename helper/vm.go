package helper

import (
	"fmt"
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

func Get_VM_List() []string {
	data, err := Read_Json(vars.CONFIG_PATH + "/config.json")
	if err != nil {
		return []string{}
	}
	vm_list, err := InterfaceSliceToStringSlice(data["vm_list"].([]interface{}))
	if err != nil {
		return []string{}
	}
	return vm_list
}

func GET_VM_UUID(vm_name string) string {
	data, err := Read_Json(vars.CONFIG_PATH + "/config.json")
	if err != nil {
		return ""
	}
	vm_uuids, err := InterfaceMapToStringMap(data["vm_uuid"].(map[string]interface{}))
	if err != nil {
		return ""
	}
	return vm_uuids[vm_name]
}

func Add_VM_To_List(vm_name string, vm_uuid string) bool {

	// add vm name
	data, err := Read_Json(vars.CONFIG_PATH + "/config.json")
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
	Write_Json(vars.CONFIG_PATH+"/config.json", data)
	return true
}

func GET_VM_Info(vm_uuid string) map[string]string {
	data, err := Read_Json(vars.CONFIG_PATH + "/" + vm_uuid + ".json")
	if err != nil {
		return nil
	}
	config, err := InterfaceMapToStringMap(data)
	if err != nil {
		return nil
	}
	return config
}
