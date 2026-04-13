package main

import "fmt"

func main() {
	for {
		var numero int
		fmt.Print("Digite um número (<=0 para sair): ")
		fmt.Scanln(&numero)

		if numero <= 0 {
			break
		}

		ehQuadradoPerfeito := false
		for i := 1; i*i <= numero; i++ {
			if i*i == numero {
				ehQuadradoPerfeito = true
				break
			}
		}

		if ehQuadradoPerfeito {
			fmt.Printf("%d é um quadrado perfeito\n", numero)
		} else {
			fmt.Printf("%d NÃO é um quadrado perfeito\n", numero)
		}
	}
}
