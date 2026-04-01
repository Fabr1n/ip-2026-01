package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var delta, x1, x2 float64
	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&b)

	fmt.Print("Digite o valor de C: ")
	fmt.Scanln(&c)
	delta = b*b - 4*a*c
	if a == 0 {
		fmt.Println("O valor de A deve ser diferente de zero para formar uma equação do segundo grau.")
		return
	}
	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
		return
	}

	if delta == 0 {
		x1 = -b / (2 * a)
		fmt.Println("RAIZ ÚNICA")
		fmt.Printf("Raiz: %.4f\n", x1)
		return
	}
	x1 = (-b + math.Sqrt(delta)) / (2 * a)
	x2 = (-b - math.Sqrt(delta)) / (2 * a)

	fmt.Println("RAÍZES DISTINTAS")
	fmt.Printf("Raiz 1: %.4f\n", x1)
	fmt.Printf("Raiz 2: %.4f\n", x2)
}

