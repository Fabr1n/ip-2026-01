package main

import "fmt"

func main() {
	var compra, venda float64
	fmt.Print("Digite o valor da compra: R$ ")
	fmt.Scanln(&compra)
	if compra < 0 {
		fmt.Println("Valor de compra inválido.")
		return
	}

	if compra < 10 {
		venda = compra * 1.70
	} else if compra < 30 {
		venda = compra * 1.50
	} else if compra < 50 {
		venda = compra * 1.40
	} else {
		venda = compra * 1.30
	}
	fmt.Printf("Valor de venda: R$ %.2f\n", venda)
}

