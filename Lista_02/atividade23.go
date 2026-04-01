package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite a idade: ")
	fmt.Scanln(&idade)
	if idade < 16 {
		fmt.Println("Classe eleitoral: Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Classe eleitoral: Eleitor obrigatório")
	} else {
		fmt.Println("Classe eleitoral: Eleitor facultativo")
	}
}

