package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scanln(&b)

	soma := a + b
	var resultado int

	if soma > 20 {
		resultado = soma + 8
	} else {
		resultado = soma - 5
	}

	fmt.Printf("Resultado: %d\n", resultado)
}
