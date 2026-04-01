package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64
	fmt.Print("Digite um número real: ")
	fmt.Scanln(&numero)
	if numero >= 0 {
		raiz := math.Sqrt(numero)
		fmt.Printf("Número informado: %.2f\n", numero)
		fmt.Printf("Raiz quadrada: %.4f\n", raiz)
	} else {
		quadrado := numero * numero
		fmt.Printf("Número informado: %.2f\n", numero)
		fmt.Printf("Como é negativo, o quadrado é: %.2f\n", quadrado)
	}
}

