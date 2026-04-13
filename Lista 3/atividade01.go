package main

import "fmt"

func main() {
	var base, expoente int

	fmt.Print("Digite a base: ")
	fmt.Scanln(&base)

	fmt.Print("Digite o expoente: ")
	fmt.Scanln(&expoente)

	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}

	fmt.Printf("%d elevado a %d = %d\n", base, expoente, resultado)
}
