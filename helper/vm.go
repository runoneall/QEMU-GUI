package helper

import "fmt"

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
