package main

import "fmt"

func main() {
	var x, fx float64
	fmt.Print("Digite o valor de x: ")
	fmt.Scanln(&x)
	if x == 2 {
		fmt.Println("Não é possível calcular f(x) para x = 2, pois haverá divisão por zero.")
		return
	}

	fx = 8 / (2 - x)
	fmt.Printf("f(%.2f) = %.4f\n", x, fx)
}

