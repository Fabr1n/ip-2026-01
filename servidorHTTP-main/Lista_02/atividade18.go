package main

import "fmt"

func main() {
	var precoNormal float64
	var categoria string
	var precoFinal float64
	fmt.Print("Digite o preço normal do DVD: R$ ")
	fmt.Scanln(&precoNormal)

	fmt.Print("Digite a categoria (comum ou lançamento): ")
	fmt.Scanln(&categoria)
	switch categoria {
	case "comum":
		precoFinal = precoNormal
	case "lançamento", "lancamento":
		precoFinal = precoNormal * 1.15
	default:
		fmt.Println("Categoria inválida.")
		return
	}
	fmt.Printf("Preço final do DVD: R$ %.2f\n", precoFinal)
}

