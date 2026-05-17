package main
import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string

	fmt.Scan(&conta, &consumo, &tipo)

	var valor float64

	switch tipo {
	case "R":
		valor = 5 + consumo*0.05
	case "C":
		valor = 500 + consumo*0.25
	case "I":
		valor = 800 + consumo*0.04
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}