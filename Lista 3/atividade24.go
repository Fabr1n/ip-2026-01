package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 0.0; a <= 6.3; a += 0.1 {
		senA := a - (math.Pow(a, 3))/6 + (math.Pow(a, 5))/120 - (math.Pow(a, 7))/5040
		fmt.Printf("Ângulo: %.1f rad, Sen(A) = %.6f\n", a, senA)
	}
}
