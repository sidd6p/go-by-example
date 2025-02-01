package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Define a constant string containing Thai characters.
	const s = "สวัสดี" // "สวัสดี" means "Hello" in Thai.

	// Print the length of the string in bytes.
	// Note: This is not the same as the number of characters (runes) because
	// some characters in UTF-8 encoding occupy more than one byte.
	fmt.Println("Len:", len(s)) // Prints: Len: 18

	// Iterate over the bytes of the string and print each byte in hexadecimal.
	// This demonstrates that each rune may be represented by multiple bytes.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // Prints the raw byte values.
	}
	fmt.Println()

	// Count the number of runes (characters) in the string.
	// utf8.RuneCountInString calculates the actual number of Unicode characters (runes).
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Prints: Rune count: 6

	// Iterate over the string using a range loop.
	// The range loop decodes runes automatically and provides both the index (byte offset)
	// and the rune value.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // Prints each rune and its starting byte position.
	}

	fmt.Println("\nUsing DecodeRuneInString")
	// Decode runes manually using `utf8.DecodeRuneInString`.
	// This demonstrates how to handle strings byte-by-byte and decode runes explicitly.
	for i, w := 0, 0; i < len(s); i += w {
		// Decode a single rune starting from the current byte index `i`.
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i) // Prints each rune and its starting byte position.
		w = width                                      // Update `w` to the width (number of bytes) of the decoded rune.

		// Pass the rune to a helper function for examination.
		examineRune(runeValue)
	}
}

// `examineRune` examines a rune and prints a message if it matches specific values.
func examineRune(r rune) {
	// Check for specific runes and print corresponding messages.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua") // Thai character 'ส'.
	}
}
