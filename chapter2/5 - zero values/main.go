package main

import "fmt"

// Zero values for various types
var (
	// Integers (int)
	intVal int

	// Floats (float64)
	floatVal float64

	// Boolean
	boolVal bool

	// String
	strVal string
)

func main() {
	fmt.Println("Zero values in Go:")

	// Integers
	fmt.Printf("int: %v\n", intVal)

	// Floats
	fmt.Printf("float64: %v\n", floatVal)

	// Boolean
	fmt.Printf("bool: %v\n", boolVal)

	// Strings
	fmt.Printf("string: '%v'\n", strVal)
}
