package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)
	if numero%5 == 0 {
		fmt.Printf("%d é divisível por 5.\n", numero)
	} else {
		fmt.Printf("%d não é divisível por 5.\n", numero)
	}
}

