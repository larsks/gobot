package tools

import (
	"fmt"
)

func ExampleReverseMap() {
	forwardMap := map[int]string{
		0:  "alice",
		3:  "bob",
		11: "mallory",
	}

	for k, v := range ReverseMap(forwardMap) {
		fmt.Printf("%s had key %d", k, v)
	}

	// Prints:
	// alice had key 0
	// bob had key 3
	// mallory had key 11
}
