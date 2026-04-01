package main

import "fmt"

func main() {
	var precoEtiqueta float64
	var codigo int
	var precoFinal float64
	fmt.Print("Digite o preço normal de etiqueta: R$ ")
	fmt.Scanln(&precoEtiqueta)

	fmt.Print("Digite o código de pagamento: ")
	fmt.Scanln(&codigo)
	switch codigo {
	case 1:
		precoFinal = precoEtiqueta * 0.90
	case 2:
		precoFinal = precoEtiqueta * 0.95
	case 3:
		precoFinal = precoEtiqueta
	case 4:
		precoFinal = precoEtiqueta * 1.10
	default:
		fmt.Println("Código de pagamento inválido.")
		return
	}
	fmt.Printf("Valor a pagar: R$ %.2f\n", precoFinal)
}

