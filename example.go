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
			"street": "Main St",
			"city":   "New York",
			"state":  "NY",
		},
	}
	if err := map2json.Convert(m, "data.json"); err != nil {
		fmt.Println(err)
	}
}
