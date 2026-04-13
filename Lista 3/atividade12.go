package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Digite um número real X: ")
	fmt.Scanln(&x)

	soma := 0.0
	termo := x

	for i := 1; i <= 20; i++ {
		fatorial := 1.0
		for j := 1; j <= i; j++ {
			fatorial *= float64(j)
		}

		if i%2 == 1 {
			soma += termo / fatorial
		} else {
			soma -= termo / fatorial
		}

		termo *= x
	}

	fmt.Printf("S = %.6f\n", soma)
}
