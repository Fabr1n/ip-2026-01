package main

import "fmt"

func main() {
	var decimal int

	fmt.Print("Digite um número inteiro positivo em base 10: ")
	fmt.Scanln(&decimal)

	binario := ""

	if decimal == 0 {
		binario = "0"
	} else {
		for decimal > 0 {
			if decimal%2 == 0 {
				binario = "0" + binario
			} else {
				binario = "1" + binario
			}
			decimal /= 2
		}
	}

	fmt.Printf("Número em binário (base 2): %s\n", binario)
}
