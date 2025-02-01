package main

import (
	"fmt"
	"reflect"
)

// The '...' before the type indicates a variadic parameter, allowing it to accept any number of arguments.
// 'nums' will be treated as a slice of int inside the function.
func sum(nums ...int) int {
	fmt.Println(reflect.TypeOf(nums)) // Output: []int (a slice of int)

	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
