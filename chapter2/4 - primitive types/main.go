package main

import "fmt"

// golang is a statically typed language, meaning the type of a variable cannot be changed

// Integers
var x int = 10

// Floats
var y float64 = 3.14

// Strings
var s string = "Hello, Go!"

// Booleans
var b bool = true

func main() {
	// Integers
	fmt.Printf("x int: %v %T\n", x, x)

	// Floats
	fmt.Printf("y float64: %v %T\n", y, y)

	// Strings
	fmt.Printf("s string: %v %T\n", s, s)

	// Booleans
	fmt.Printf("b bool: %v %T\n", b, b)
}
