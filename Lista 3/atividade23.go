package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite o número de termos (N): ")
	_, err := fmt.Scan(&n)

	if err != nil || n <= 0 {
		fmt.Println("Valor inválido! Por favor, insira um número inteiro maior que zero.")
		return
	}

	soma := 0.0
	numerador := 1000.0
	denominador := 1.0
	sinal := 1.0

	for i := 1; i <= n; i++ {
		termo := sinal * (numerador / denominador)
		soma += termo

		numerador -= 3.0
		denominador += 1.0
		sinal *= -1.0
	}

	fmt.Printf("O resultado da série considerando os %d primeiros termos é: %.4f\n", n, soma)
}
