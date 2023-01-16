package main

import (
	"fmt"

	"github.com/johnesleyer/map2json/map2json"
)

func main() {
	m := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "Anytown",
			"state":  "XX",
			"zip":    "12345",
		},
	}

	// convert map to json and write to a file
	if err := map2json.ToJson(m, "example.json"); err != nil {
		fmt.Println("Error while converting to json:", err)
	}

	// convert json file to map
	result, err := map2json.ToMap("example.json")
	if err != nil {
		fmt.Println("Error while converting to map:", err)
	}
	fmt.Println(result)
}
