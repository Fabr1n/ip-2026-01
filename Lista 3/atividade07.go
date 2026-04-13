package main

import "fmt"

func main() {
	var numero float64
	soma := 0.0
	quantidade := 0
	maior := -999999.0
	menor := 999999.0
	somaPares := 0.0
	quantidadePares := 0
	quantidadeImpares := 0

	for {
		fmt.Print("Digite um número (30.000 para finalizar): ")
		fmt.Scanln(&numero)

		if numero == 30.000 {
			break
		}

		soma += numero
		quantidade++

		if numero > maior {
			maior = numero
		}
		if numero < menor {
			menor = numero
		}

		intNum := int(numero)
		if intNum%2 == 0 {
			somaPares += numero
			quantidadePares++
		} else {
			quantidadeImpares++
		}
	}

	media := soma / float64(quantidade)
	percentualImpares := (float64(quantidadeImpares) / float64(quantidade)) * 100

	fmt.Printf("\na) Soma: %.2f\n", soma)
	fmt.Printf("b) Quantidade de números: %d\n", quantidade)
	fmt.Printf("c) Média: %.2f\n", media)
	fmt.Printf("d) Maior número: %.2f\n", maior)
	fmt.Printf("e) Menor número: %.2f\n", menor)

	if quantidadePares > 0 {
		mediaPares := somaPares / float64(quantidadePares)
		fmt.Printf("f) Média dos pares: %.2f\n", mediaPares)
	} else {
		fmt.Println("f) Nenhum número par digitado")
	}

	fmt.Printf("g) Percentagem de ímpares: %.2f%%\n", percentualImpares)
}
