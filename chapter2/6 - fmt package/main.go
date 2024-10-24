package main

import "fmt"

func main() {
	// Print sem formatação
	fmt.Print("Hello, ")  // Print sem newline
	fmt.Println("World!") // Print com newline
	fmt.Println("Go is awesome")

	// Printf com formatação
	name := "John"
	age := 30
	height := 1.85
	isStudent := false

	fmt.Printf("Name: %s, Age: %d, Height: %.2f, Student: %t\n", name, age, height, isStudent)

	// Especificadores de formato comuns:
	// %v -> Valor padrão
	// %T -> Tipo do valor
	// %d -> Inteiro
	// %f -> Float
	// %s -> String
	// %t -> Booleano

	// Usando %v (valor padrão) e %T (tipo)
	fmt.Printf("Default format: %v\n", name) // Valor padrão da string
	fmt.Printf("Type of name: %T\n", name)   // Tipo de variável

	fmt.Printf("Default format: %v\n", age) // Valor padrão de int
	fmt.Printf("Type of age: %T\n", age)    // Tipo de variável

	// Formatando floats com precisão
	fmt.Printf("Height (2 decimal places): %.2f\n", height) // Exibe com 2 casas decimais

	// Booleano
	fmt.Printf("Is student? %t\n", isStudent)

	// Exibindo números em diferentes bases
	num := 255
	fmt.Printf("Decimal: %d\n", num)     // Decimal
	fmt.Printf("Binary: %b\n", num)      // Binário
	fmt.Printf("Hexadecimal: %x\n", num) // Hexadecimal

	// Sprintf: Formata e retorna uma string (não imprime no console)
	formattedString := fmt.Sprintf("Formatted string: Name=%s, Age=%d", name, age)
	fmt.Println(formattedString)

	// Errorf: Formata e cria um erro
	err := fmt.Errorf("an error occurred: %s", "connection lost")
	fmt.Println("Error:", err)

	// Lendo a entrada do usuário
	var userInput string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userInput) // Lê uma linha de input do usuário
	fmt.Printf("Hello, %s!\n", userInput)
}
