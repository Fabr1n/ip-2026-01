package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Digite o lado A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o lado B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o lado C: ")
	fmt.Scan(&c)

	if !tri(a, b, c) {
		fmt.Println("Os valores não formam um triângulo.")
		return
	}

	if a == b && b == c {
		fmt.Println("Triângulo equilátero")
	} else if a == b || b == c || a == c {
		fmt.Println("Triângulo isósceles")
	} else {
		fmt.Println("Triângulo escaleno")
	}
}

func tri(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}