package main

import "fmt"

func MEDIA(X, Y, Z float64) float64 {
	return (X + Y + Z) / 3
}

func main() {
	var x, y, z float64

	fmt.Print("Digite o primeiro valor real: ")
	fmt.Scanln(&x)

	fmt.Print("Digite o segundo valor real: ")
	fmt.Scanln(&y)

	fmt.Print("Digite o terceiro valor real: ")
	fmt.Scanln(&z)

	media := MEDIA(x, y, z)

	fmt.Printf("A média de %.2f, %.2f e %.2f é: %.2f\n", x, y, z, media)
}

