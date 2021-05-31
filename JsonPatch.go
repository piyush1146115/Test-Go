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

	//Comparing JSON documents
	original1 := []byte(`{"name": "John", "age": 24, "height": 3.21}`)
	similar := []byte(`
		{
			"age": 24,
			"height": 3.21,
			"name": "John"
		}
	`)
	different := []byte(`{"name": "Jane", "age": 20, "height": 3.37}`)

	if jsonpatch.Equal(original1, similar) {
		fmt.Println(`"original1" is structurally equal to "similar"`)
	}

	if !jsonpatch.Equal(original1, different) {
		fmt.Println(`"original1" is _not_ structurally equal to "different"`)
	}





	//Combine merge patches
	nameAndHeight := []byte(`{"height":null,"name":"Jane"}`)
	ageAndEyes := []byte(`{"age":4.23,"eyes":"blue"}`)
	// Let's combine these merge patch documents...
	combinedPatch, err := jsonpatch.MergeMergePatches(nameAndHeight, ageAndEyes)
	if err != nil {
		panic(err)
	}

	// Apply each patch individual against the original document
	withoutCombinedPatch, err := jsonpatch.MergePatch(original, nameAndHeight)
	if err != nil {
		panic(err)
	}

	withoutCombinedPatch, err = jsonpatch.MergePatch(withoutCombinedPatch, ageAndEyes)
	if err != nil {
		panic(err)
	}

	// Apply the combined patch against the original document

	withCombinedPatch, err := jsonpatch.MergePatch(original, combinedPatch)
	if err != nil {
		panic(err)
	}

	// Do both result in the same thing? They should!
	if jsonpatch.Equal(withCombinedPatch, withoutCombinedPatch) {
		fmt.Println("Both JSON documents are structurally the same!")
	}

	fmt.Printf("combined merge patch: %s", combinedPatch)
}