package main

import "fmt"

func main() {
	var a, b, c int
	var menor, inter, maior int
	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&b)

	fmt.Print("Digite o valor de C: ")
	fmt.Scanln(&c)
	if a < b && a < c {
		menor = a
		if b < c {
			inter = b
			maior = c
		} else {
			inter = c
			maior = b
		}
	} else if b < a && b < c {
		menor = b
		if a < c {
			inter = a
			maior = c
		} else {
			inter = c
			maior = a
		}
	} else {
		menor = c
		if a < b {
			inter = a
			maior = b
		} else {
			inter = b
			maior = a
		}
	}
	fmt.Printf("Ordem crescente: %d, %d, %d\n", menor, inter, maior)
}

