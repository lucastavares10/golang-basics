package main

import (
	"fmt"
)

func main() {
	// string normal e uma raw string literal
	str := "Hello, Go!"

	fmt.Printf("string: %v, Tipo: %T\n", str, str)

	// raw string literal
	rawStr := `Isto é um 
raw string literal`

	fmt.Printf("raw string literal: %v, Tipo: %T\n", rawStr, rawStr)

	// slice de bytes
	byteSlice := []byte(str)
	fmt.Printf("slice de bytes: %v, tipo: %T\n", byteSlice, byteSlice)

	// str[0] = 'h' // "cannot assign to str[0]"

	// any byte
	fmt.Println("bytes da string:")
	for i, b := range byteSlice {
		fmt.Printf("índice: %d, byte: %d, char: %c\n", i, b, b)
	}
}
