package main
import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	fmt.Scan(&h, &a)

	base := (3 * a * a * math.Sqrt(3)) / 2
	v := (1.0 / 3.0) * base * h

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", v)
}