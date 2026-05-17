package main
import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta := b*b - 4*a*c
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}