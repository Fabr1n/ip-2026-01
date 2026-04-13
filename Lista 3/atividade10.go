package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scanln(&n)

	if n < 3 {
		fmt.Println("N deve ser maior ou igual a 3")
		return
	}

	var a, b int = 0, 1

	fmt.Print("Sequência de Fibonacci: ")
	fmt.Print(a, " ")
	fmt.Print(b, " ")

	for i := 3; i <= n; i++ {
		c := a + b
		fmt.Print(c, " ")
		a = b
		b = c
	}

	fmt.Println()
}
