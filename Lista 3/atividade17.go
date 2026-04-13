package main

import "fmt"

func main() {
	fmt.Println("Índices de todos os elementos de uma matriz 10x10:")

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("[%d][%d] ", i, j)
		}
		fmt.Println()
	}
}
