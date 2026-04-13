package main

import "fmt"

func main() {
	var n, a1, a2 int

	fmt.Print("Digite o primeiro termo: ")
	fmt.Scanln(&a1)

	fmt.Print("Digite o segundo termo: ")
	fmt.Scanln(&a2)

	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scanln(&n)

	if n < 3 {
		fmt.Println("N deve ser maior ou igual a 3")
		return
	}

	fmt.Print("Série de Fetuccine: ")
	fmt.Print(a1, " ", a2, " ")

	for i := 3; i <= n; i++ {
		var novoTermo int

		if i%2 == 1 {
			novoTermo = a2 + a1
		} else {
			novoTermo = a2 - a1
		}

		fmt.Print(novoTermo, " ")
		a1 = a2
		a2 = novoTermo
	}

	fmt.Println()
}
