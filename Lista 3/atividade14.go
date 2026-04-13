package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scanln(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scanln(&n2)

	fmt.Printf("Números primos entre %d e %d:\n", n1, n2)

	for num := n1; num <= n2; num++ {
		if isPrimo(num) {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}

func isPrimo(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
