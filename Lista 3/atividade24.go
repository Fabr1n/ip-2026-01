package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ângulo A (rad) | Seno(A) Aproximado")
	fmt.Println("-----------------------------------")

	for i := 0; i <= 63; i++ {
		A := float64(i) / 10.0
		termo1 := A
		termo2 := math.Pow(A, 3) / 6.0
		termo3 := math.Pow(A, 5) / 120.0
		termo4 := math.Pow(A, 7) / 5040.0

		senA := termo1 - termo2 + termo3 - termo4
		fmt.Printf("%14.1f | %18.6f\n", A, senA)
	}
}
