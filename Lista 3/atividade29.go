package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&n)

	soma := 0

	for i := 1; i <= n; i++ {
		soma += i
	}

	fmt.Printf("Somatório de 1 a %d = %d\n", n, soma)
}
