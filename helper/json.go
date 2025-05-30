package helper

import (
	"encoding/json"
)

func WriteJson(file_path string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return WriteFile(file_path, jsonData)
}

func ReadJson(file_path string) (map[string]interface{}, error) {
	content, err := ReadFile(file_path)
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
