package main

import "fmt"

func main() {
	h := 0.0

	for i := 1; i <= 50; i++ {
		numerador := float64(2*i - 1)
		denominador := float64(i)
		h += numerador / denominador
	}

	fmt.Printf("H = %.6f\n", h)
}
