package main

import "fmt"

func main() {
	var idade int
	var classificacao string
	fmt.Print("Digite a idade: ")
	fmt.Scanln(&idade)
	if idade <= 2 {
		classificacao = "Recém-nascido"
	} else if idade <= 11 {
		classificacao = "Criança"
	} else if idade <= 19 {
		classificacao = "Adolescente"
	} else if idade <= 55 {
		classificacao = "Adulto"
	} else {
		classificacao = "Idoso"
	}
	fmt.Printf("Classificação: %s\n", classificacao)
}

