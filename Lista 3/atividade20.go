package main

import "fmt"

func main() {
	fmt.Println("Índices dos elementos ABAIXO da diagonal principal de uma matriz 10x10:")

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("[%d][%d] ", i, j)
		}
	}
	fmt.Println()
}
