package main

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)
func main() {
	original := []byte(`{"name": "John", "age": 24, "height": 3.21}`)
	patchJSON := []byte(`[
		{"op": "replace", "path": "/name", "value": "Jane"},
		{"op": "remove", "path": "/height"}
	]`)

	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		panic(err)
	}

	modified, err := patch.Apply(original)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original document: %s\n", original)
	fmt.Printf("Modified document: %s\n", modified)
}