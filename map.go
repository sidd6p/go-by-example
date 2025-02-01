package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"k1": 10,
		"k2": 20,
		"k3": 30,
	}

	// Check if key "k2" exists in the map. The second return value is a boolean indicating presence.
	_, prs := m["k2"]

	if prs {
		fmt.Println("Key 'k2' exists in the map.")
	} else {
		fmt.Println("Key 'k2' does not exist in the map.")
	}

	delete(m, "k2")
	fmt.Println("map:", m)
}
