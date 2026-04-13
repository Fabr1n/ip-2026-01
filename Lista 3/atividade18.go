package main

import "fmt"

func main() {
	fmt.Println("Índices da diagonal principal de uma matriz 10x10:")

	for i := 0; i < 10; i++ {
		fmt.Printf("[%d][%d] ", i, i)
	}
	fmt.Println()
}
