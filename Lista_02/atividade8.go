package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)
	if numero > 20 && numero < 90 {
		fmt.Printf("%d está compreendido entre 20 e 90.\n", numero)
	} else {
		fmt.Printf("%d não está compreendido entre 20 e 90.\n", numero)
	}
}

