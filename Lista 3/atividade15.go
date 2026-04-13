package main

import "fmt"

func main() {
	var n int

	fmt.Print("Digite quantos termos deseja: ")
	fmt.Scanln(&n)

	fmt.Println("Série dos quadrados:")
	for i := 1; i <= n; i++ {
		quadrado := i * i
		fmt.Print(quadrado, " ")
	}
	fmt.Println()
}
