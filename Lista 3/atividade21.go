package main

import "fmt"

func main() {
	var b, n int

	fmt.Print("Digite o valor de b (>= 2): ")
	fmt.Scanln(&b)

	fmt.Print("Digite o valor de n (> 1): ")
	fmt.Scanln(&n)

	resultado := 1

	for i := 0; i < n; i++ {
		resultado *= b
	}

	fmt.Printf("%d elevado a %d = %d\n", b, n, resultado)
}
