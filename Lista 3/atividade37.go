package main

import "fmt"

func main() {
	var octal int

	fmt.Print("Digite um número inteiro positivo em base 8: ")
	fmt.Scanln(&octal)

	decimal := 0
	potencia := 1

	for octal > 0 {
		digito := octal % 10
		decimal += digito * potencia
		potencia *= 8
		octal /= 10
	}

	fmt.Printf("Número em decimal (base 10): %d\n", decimal)
}
