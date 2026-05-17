package main
import "fmt"

func main() {
	var s float64
	fmt.Scan(&s)

	if s <= 300 {
		s *= 1.5
	} else {
		s *= 1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", s)
}