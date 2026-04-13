package main

import "fmt"

func main() {
	graos := 1
	totalGraos := int64(0)

	fmt.Println("Tabuleiro de xadrez - Número de grãos por quadro:")

	for quadro := 1; quadro <= 64; quadro++ {
		totalGraos += int64(graos)

		if quadro%8 == 0 {
			fmt.Printf("Quadro %d: %d grãos | Total até aqui: %d\n", quadro, graos, totalGraos)
		}

		graos *= 2
	}

	fmt.Printf("\nTotal de grãos no tabuleiro: %d\n", totalGraos)
}
