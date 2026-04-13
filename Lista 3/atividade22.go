package main

import "fmt"

func main() {
	soma := 0.0

	for i := 1; i <= 37; i++ {
		numerador := float64((38 - i) * (39 - i))
		denominador := float64(i)
		soma += numerador / denominador
	}

	fmt.Printf("S = %.6f\n", soma)
}
