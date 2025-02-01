package main

import "fmt"

// A function that returns two integer values
func vals() (int, int) {
	return 3, 7 // Returns two values: 3 and 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
