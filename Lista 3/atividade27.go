package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Print("Digite o valor de x em radianos: ")
	fmt.Scanln(&x)

	cossenoCalculado := 1.0

	for i := 1; i <= 10; i++ {
		numerador := math.Pow(x, float64(2*i))

		fatorial := 1.0
		for j := 1; j <= 2*i; j++ {
			fatorial *= float64(j)
		}

		termo := numerador / fatorial

		if i%2 == 1 {
			cossenoCalculado -= termo
		} else {
			cossenoCalculado += termo
		}
	}

	cossenoReal := math.Cos(x)
	diferenca := cossenoCalculado - cossenoReal

	fmt.Printf("a) Cosseno calculado: %.6f\n", cossenoCalculado)
	fmt.Printf("b) Cosseno (função): %.6f\n", cossenoReal)
	fmt.Printf("c) Diferença: %.6f\n", diferenca)
}
