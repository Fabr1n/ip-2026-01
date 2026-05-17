package main
import (
	"fmt"
	"math"
)

func main() {
	var r, h float64
	fmt.Scan(&r)
	fmt.Scan(&h)

	area := 2*math.Pi*r*r + 2*math.Pi*r*h
	custo := area * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}