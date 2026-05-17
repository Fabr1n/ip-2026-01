package main

import "fmt"

func main() {
	var tipo string
	var consumo float64
	var conta float64
	fmt.Print("Digite o tipo de consumidor (residencial, comercial ou industrial): ")
	fmt.Scanln(&tipo)

	fmt.Print("Digite o consumo em m3: ")
	fmt.Scanln(&consumo)
	switch tipo {
	case "residencial":
		conta = 5 + (0.05 * consumo)
	case "comercial":
		if consumo <= 80 {
			conta = 500
		} else {
			conta = 500 + ((consumo - 80) * 0.25)
		}
	case "industrial":
		if consumo <= 100 {
			conta = 800
		} else {
			conta = 800 + ((consumo - 100) * 0.04)
		}
	default:
		fmt.Println("Tipo de consumidor inválido.")
		return
	}
	fmt.Printf("Valor da conta: R$ %.2f\n", conta)
}

