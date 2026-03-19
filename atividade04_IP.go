package main
import "fmt"

func main() {
	var salario, kw float64
	fmt.Scan(&salario, &kw)

	precoKW := (0.7 * salario) / 100
	total := precoKW * kw
	desconto := total * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", precoKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", total)
	fmt.Printf("Custo com desconto: R$ %.2f\n", desconto)
}