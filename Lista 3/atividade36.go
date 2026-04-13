package main

import "fmt"

func main() {
	var decimal int

	fmt.Print("Digite um número inteiro positivo em base 10: ")
	fmt.Scanln(&decimal)

	hexadecimal := ""
	digitos := "0123456789ABCDEF"

	if decimal == 0 {
		hexadecimal = "0"
	} else {
		for decimal > 0 {
			hexadecimal = string(digitos[decimal%16]) + hexadecimal
			decimal /= 16
		}
	}

	fmt.Printf("Número em hexadecimal (base 16): %s\n", hexadecimal)
}
