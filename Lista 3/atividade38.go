package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cpf string

	fmt.Print("Digite o CPF (9 dígitos): ")
	fmt.Scanln(&cpf)

	if len(cpf) != 9 {
		fmt.Println("CPF inválido!")
		return
	}

	soma1 := 0
	for i := 0; i < 9; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma1 += digito * (10 - i)
	}

	resto1 := soma1 % 11
	var digito1 int

	if resto1 < 2 {
		digito1 = 0
	} else {
		digito1 = 11 - resto1
	}

	soma2 := 0
	for i := 0; i < 9; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma2 += digito * (11 - i)
	}

	soma2 += digito1 * 2

	resto2 := soma2 % 11
	var digito2 int

	if resto2 < 2 {
		digito2 = 0
	} else {
		digito2 = 11 - resto2
	}

	cpfCompleto := cpf + strconv.Itoa(digito1) + strconv.Itoa(digito2)

	fmt.Printf("CPF com dígitos verificadores: %s.%s.%s-%s%s\n", 
		cpfCompleto[0:3], cpfCompleto[3:6], cpfCompleto[6:9], 
		string(rune(digito1)+'0'), string(rune(digito2)+'0'))

	var verificacao string
	fmt.Print("Digite os 2 dígitos verificadores: ")
	fmt.Scanln(&verificacao)

	if verificacao[0:1] == strconv.Itoa(digito1) && verificacao[1:2] == strconv.Itoa(digito2) {
		fmt.Println("CPF válido!")
	} else {
		fmt.Println("CPF inválido!")
	}
}
