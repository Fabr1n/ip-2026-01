package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64
	var volume, area float64
	fmt.Println("1 - Cone reto")
	fmt.Println("2 - Cilindro")
	fmt.Println("3 - Esfera")
	fmt.Print("Escolha a opção: ")
	fmt.Scanln(&opcao)
	switch opcao {
	case 1:
		fmt.Print("Digite o raio do cone: ")
		fmt.Scanln(&raio)
		fmt.Print("Digite a altura do cone: ")
		fmt.Scanln(&altura)

		volume = (math.Pi * raio * raio * altura) / 3
		area = math.Pi * raio * math.Sqrt(raio*raio+altura*altura)

	case 2:
		fmt.Print("Digite o raio do cilindro: ")
		fmt.Scanln(&raio)
		fmt.Print("Digite a altura do cilindro: ")
		fmt.Scanln(&altura)

		volume = math.Pi * raio * raio * altura
		area = 2 * math.Pi * raio * altura

	case 3:
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scanln(&raio)

		volume = (4 * math.Pi * math.Pow(raio, 3)) / 3
		area = 4 * math.Pi * raio * raio

	default:
		fmt.Println("Opção inválida.")
		return
	}
	fmt.Printf("Volume: %.4f\n", volume)
	fmt.Printf("Área da superfície: %.4f\n", area)
}

