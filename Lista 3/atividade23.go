package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite a quantidade de termos: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Erro: N deve ser maior que 0!")
		return
	}

	soma := 0.0

	for i := 1; i <= n; i++ {
		var termo float64

		if i%2 == 1 {
			numerador := 1000.0 - float64((i-1)*3)
			termo = numerador / float64(i)
		} else {
			numerador := 1000.0 - float64((i-1)*3)
			termo = -(numerador * numerador) / float64(i)
		}

		soma += termo
	}

	fmt.Printf("Soma = %.6f\n", soma)
}
