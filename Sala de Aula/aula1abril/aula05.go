package main

import "fmt"

func Fatorial(n int) int {
	resultado := 1

	for i := 2; i <= n; i++ {
		resultado *= i
	}

	return resultado
}

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	fatorial := Fatorial(numero)

	fmt.Printf("O fatorial de %d é: %d\n", numero, fatorial)
}
