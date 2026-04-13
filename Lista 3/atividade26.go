package main

import "fmt"

func main() {
	soma := 0.0

	for i := 0; i < 20; i++ {
		numerador := float64(100 - i)
		fatorial := 1.0

		for j := 1; j <= i; j++ {
			fatorial *= float64(j)
		}

		soma += numerador / fatorial
	}

	fmt.Printf("Soma dos 20 primeiros termos: %.6f\n", soma)
}
