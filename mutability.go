package main

import "fmt"

func main() {
	// Strings are immutable, so we cannot modify them directly.
	str := "hello"
	// str[0] = 'H' // ❌ Compilation error: cannot assign to str[0]
	fmt.Println("Original String:", str)

	// ✅ Method 1: Using Rune Slice (Handles Unicode characters)
	runes := []rune(str)                                  // Convert string to a rune slice
	runes[0] = 'H'                                        // Modify the first character
	newStr := string(runes)                               // Convert back to a string
	fmt.Println("Modified String (Using Runes):", newStr) // Output: "Hello"

	// ✅ Method 2: Using Byte Slice (Works for ASCII characters)
	bytes := []byte(str)                                       // Convert string to a byte slice
	bytes[0] = 'H'                                             // Modify the first character
	newStrBytes := string(bytes)                               // Convert back to a string
	fmt.Println("Modified String (Using Bytes):", newStrBytes) // Output: "Hello"

	// ✅ Arrays are mutable, but their size is fixed
	arr := [3]int{1, 2, 3}
	arr[0] = 10                         // Allowed, modifies the array
	fmt.Println("Modified Array:", arr) // Output: [10, 2, 3]

	// ✅ Slices are mutable and resizable
	slice := []int{1, 2, 3}
	slice[0] = 10                         // Allowed, modifies the slice
	slice = append(slice, 4)              // Resizes the slice dynamically
	fmt.Println("Modified Slice:", slice) // Output: [10, 2, 3, 4]
}
