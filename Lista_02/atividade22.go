package main

import "fmt"

func main() {
	var matricula int
	var horasExtras int
	var salarioMinimo float64 = 788.00
	var valorHoraExtra float64 = 10.00
	var salarioHoraExtra float64
	var salarioBruto float64
	var inss float64
	var ir float64
	var salarioLiquido float64
	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scanln(&matricula)

	fmt.Print("Digite a quantidade de horas extras trabalhadas: ")
	fmt.Scanln(&horasExtras)
	salarioHoraExtra = float64(horasExtras) * valorHoraExtra
	salarioBruto = 3*salarioMinimo + salarioHoraExtra
	if salarioBruto > 1500 {
		inss = salarioBruto * 0.12
	}

	if salarioBruto > 2000 {
		ir = salarioBruto * 0.20
	}

	salarioLiquido = salarioBruto - inss - ir
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Salário hora-extra: R$ %.2f\n", salarioHoraExtra)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", inss)
	fmt.Printf("Desconto IR: R$ %.2f\n", ir)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}

