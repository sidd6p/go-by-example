package main

import "fmt"

func main() {
	// Using range with arrays
	nums := [5]int{1, 2, 3, 4, 5}
	sum := 0
	// 'range' returns two values: index and value, but we ignore the index using '_'
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// Using range with maps
	mps := map[int]string{1: "one", 2: "two"}
	// 'range' returns key-value pairs from the map
	for key, val := range mps {
		fmt.Println(key, "==>", val)
	}

	// Using range with strings
	name := "Sidd"
	// 'range' iterates over Unicode code points (runes) of the string
	// It returns the index and the Unicode value (rune) for each character in the string
	for idx, char := range name {
		fmt.Printf("Index: %d, Unicode: %U, Character: %c\n", idx, char, char) // Prints the index, Unicode, and character
	}
}
