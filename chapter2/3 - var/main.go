package main

import "fmt"

var x = 10

func main() {
	teste(x)
}

func teste(x int) {
	fmt.Printf("teste function: %v\n", x)
}
