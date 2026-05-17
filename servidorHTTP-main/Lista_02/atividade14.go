package main

import "fmt"

func main() {
	var precoBase float64
	var ac, pintura, vidro, direcao bool
	var precoFinal float64
	fmt.Print("Digite o preço inicial de fábrica: R$ ")
	fmt.Scanln(&precoBase)
	fmt.Print("Deseja ar condicionado? (true/false): ")
	fmt.Scanln(&ac)

	fmt.Print("Deseja pintura metálica? (true/false): ")
	fmt.Scanln(&pintura)

	fmt.Print("Deseja vidro elétrico? (true/false): ")
	fmt.Scanln(&vidro)

	fmt.Print("Deseja direção hidráulica? (true/false): ")
	fmt.Scanln(&direcao)
	precoFinal = precoBase

	if ac {
		precoFinal += 1750
	}
	if pintura {
		precoFinal += 800
	}
	if vidro {
		precoFinal += 1200
	}
	if direcao {
		precoFinal += 2000
	}
	fmt.Printf("Preço final do carro: R$ %.2f\n", precoFinal)
}

