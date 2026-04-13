package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&n1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&n2)

	resultado := 0

	for i := 0; i < n2; i++ {
		resultado += n1
	}

	fmt.Printf("%d x %d = %d\n", n1, n2, resultado)
}
