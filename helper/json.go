package helper

import (
	"encoding/json"
)

func InvertMap(m map[string]string) map[string]string {
	inverted := make(map[string]string, len(m))
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}

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
