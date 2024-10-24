package main

import "fmt"

// Criando um tipo personalizado chamado 'Age' a partir do tipo int
type Age int

func main() {
	// Usando o tipo personalizado 'Age'
	var myAge Age = 30
	fmt.Printf("My age is: %v and the type is: %T\n", myAge, myAge)

	// Convertendo 'Age' de volta para int (se necessário)
	ageInYears := int(myAge)
	fmt.Printf("My age as an int: %v and the type is: %T\n", ageInYears, ageInYears)
}
