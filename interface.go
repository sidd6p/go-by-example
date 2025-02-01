package main

import (
	"fmt"
	"math"
)

// Defining the 'geometry' interface with two method signatures: area() and perim().
type geometry interface {
	area() float64
	perim() float64
}

type react struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Implementing the 'area' method for the 'react' type
func (r react) area() float64 {
	return r.width * r.height
}

// Implementing the 'perim' method for the 'react' type
func (r react) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implementing the 'area' method for the 'circle' type
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implementing the 'perim' method for the 'circle' type
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 'measure' function accepts a 'geometry' interface and calls its methods
func measure(g geometry) {
	// Print the geometry object, its area, and perimeter
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	// Type assertion: Checking if 'g' is a 'circle' type and extracting its value
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}

	// Type switch: Checking the actual type of 'g' at runtime
	switch g.(type) {
	case circle:
		fmt.Println("circle")
	case react:
		fmt.Println("react")
	default:
		fmt.Println("Not known")
	}
}

func main() {
	r := react{width: 3, height: 10}

	c := circle{radius: 10}

	// Call the 'measure' function for both the circle and rectangle
	measure(c)
	measure(r)

	// Call the 'detectCircle' function for both the circle and rectangle
	detectCircle(c)
	detectCircle(r)
}
