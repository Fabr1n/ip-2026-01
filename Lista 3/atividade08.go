package main

import "fmt"

func main() {
	soma := 0

	fmt.Println("Números de 1 a 20:")
	for i := 1; i <= 20; i++ {
		fmt.Print(i, " ")
		soma += i
	}

	fmt.Printf("\n\nSoma de 1 a 20: %d\n", soma)
}
