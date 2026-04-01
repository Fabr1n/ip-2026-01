package main

import "fmt"

func main() {
	var destino, retorno int
	var preco float64
	fmt.Println("Destino: 1-Norte, 2-Nordeste, 3-Centro-Oeste, 4-Sul")
	fmt.Print("Digite o destino: ")
	fmt.Scanln(&destino)

	fmt.Println("Retorno: 1-Sim, 2-Não")
	fmt.Print("A viagem inclui retorno? ")
	fmt.Scanln(&retorno)
	if retorno != 1 && retorno != 2 {
		fmt.Println("Valor inválido para ida e volta.")
		return
	}

	switch destino {
	case 1:
		if retorno == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2:
		if retorno == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3:
		if retorno == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4:
		if retorno == 1 {
			preco = 550
		} else {
			preco = 300
		}
	default:
		fmt.Println("Valor inválido para destino.")
		return
	}
	fmt.Printf("Preço da passagem: R$ %.2f\n", preco)
}

