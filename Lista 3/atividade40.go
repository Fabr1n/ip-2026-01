package main

import "fmt"

func main() {
	precoBase := 6.0
	ingresosBase := 130
	despesas := 300.0
	maiorLucro := 0.0
	melhorPreco := 0.0
	melhorIngressos := 0

	fmt.Println("Preço (R$) | Ingressos | Lucro (R$)")
	fmt.Println("------------------------------------")

	for preco := 6.0; preco >= 1.0; preco -= 0.6 {
		reducao := (precoBase - preco) / 0.6
		ingressos := ingresosBase + int(reducao)*30
		receita := preco * float64(ingressos)
		lucro := receita - despesas

		fmt.Printf("%.2f       | %d      | %.2f\n", preco, ingressos, lucro)

		if lucro > maiorLucro {
			maiorLucro = lucro
			melhorPreco = preco
			melhorIngressos = ingressos
		}
	}

	fmt.Printf("\nMaior lucro: R$ %.2f\n", maiorLucro)
	fmt.Printf("Preço: R$ %.2f\n", melhorPreco)
	fmt.Printf("Ingressos vendidos: %d\n", melhorIngressos)
}
