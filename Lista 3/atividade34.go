package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("Digite N1: ")
	fmt.Scanln(&n1)

	fmt.Print("Digite N2: ")
	fmt.Scanln(&n2)

	a := n1
	b := n2

	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	mdc := a

	mmc := (n1 * n2) / mdc

	fmt.Printf("MMC(%d, %d) = %d\n", n1, n2, mmc)
}
