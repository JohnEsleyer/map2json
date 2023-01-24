package map2json

import (
	"encoding/json"
	"os"
)

//Convert map to json function
func ToJson(m map[string]interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(m)
}

//Convert json file to map
func ToMap(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result map[string]interface{}
	dec := json.NewDecoder(file)
	if err := dec.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
