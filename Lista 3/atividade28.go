package main

import (
	"fmt"
	"math"
)

func main() {
	soma := 0.0

	for i := 1; i <= 51; i += 2 {
		n1 := i
		n2 := i + 2
		n3 := i + 4

		denominador := math.Pow(float64(n1), 3) - math.Pow(float64(n2), 3) + math.Pow(float64(n3), 3)

		soma += 1.0 / denominador
	}

	pi := math.Pow(soma*32, 1.0/3.0)

	fmt.Printf("Valor de π com 51 termos: %.6f\n", pi)
}
