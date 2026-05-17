package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Digite o valor de A: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o valor de B: ")
	fmt.Scanln(&b)
	if b == 0 {
		fmt.Println("Não é possível dividir por zero.")
		return
	}

	if a%b == 0 {
		fmt.Printf("%d é divisível por %d.\n", a, b)
	} else {
		fmt.Printf("%d não é divisível por %d.\n", a, b)
	}
}

