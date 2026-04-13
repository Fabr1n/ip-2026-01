package main

import "fmt"

func main() {
	var salarioCarlos float64

	fmt.Print("Digite o salário de Carlos: R$ ")
	fmt.Scanln(&salarioCarlos)

	salarioJoao := salarioCarlos / 3
	meses := 0

	for salarioJoao < salarioCarlos {
		salarioCarlos *= 1.02
		salarioJoao *= 1.05
		meses++
	}

	fmt.Printf("Quantidade de meses para João igualar ou ultrapassar Carlos: %d meses\n", meses)
	fmt.Printf("Valor de Carlos após %d meses: R$ %.2f\n", meses, salarioCarlos)
	fmt.Printf("Valor de João após %d meses: R$ %.2f\n", meses, salarioJoao)
}
