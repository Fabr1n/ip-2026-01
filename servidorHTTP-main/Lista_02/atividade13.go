package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo de 3 casas: ")
	fmt.Scanln(&numero)
	if numero < 100 || numero > 999 {
		fmt.Println("O número informado não possui exatamente 3 casas.")
		return
	}
	dezenas := (numero / 10) % 10
	fmt.Printf("O algarismo da casa das dezenas é: %d\n", dezenas)
}

