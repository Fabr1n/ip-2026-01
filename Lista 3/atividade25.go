package main

import (
	"fmt"
	"math"
)

func main() {
	soma := 0.0

	soma += 1.0 / 225
	soma -= 2.0 / 196
	soma += 4.0 / 169
	soma -= 8.0 / 144

	for i := 4; i <= 7; i++ {
		numerador := math.Pow(2, float64(i))
		denominador := math.Pow(169-15*float64(i-3), 2)
		if i%2 == 0 {
			soma -= numerador / denominador
		} else {
			soma += numerador / denominador
		}
	}

	soma += 16384.0 / 1.0

	fmt.Printf("S = %.6f\n", soma)
}
