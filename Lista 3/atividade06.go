package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&numero)

	ehTriangular := false
	var i int

	for i := 1; i*i*i < numero; i++ {
		if i*(i+1)*(i+2) == numero {
			ehTriangular = true
			fmt.Printf("%d = %d x %d x %d\n", numero, i, i+1, i+2)
			break
		}
	}

	if !ehTriangular {
		fmt.Printf("%d NÃO é um número triangular\n", numero)
	} else {
		fmt.Printf("%d é um número triangular\n", numero)
	}
}
