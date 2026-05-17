package main

import "fmt"

func main() {
	var matricula int
	var nota1, nota2, nota3, mediaExercicios float64
	var mediaAproveitamento float64
	var conceito string
	fmt.Print("Digite o número de identificação do aluno: ")
	fmt.Scanln(&matricula)

	fmt.Print("Digite a nota 1: ")
	fmt.Scanln(&nota1)

	fmt.Print("Digite a nota 2: ")
	fmt.Scanln(&nota2)

	fmt.Print("Digite a nota 3: ")
	fmt.Scanln(&nota3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scanln(&mediaExercicios)
	mediaAproveitamento = (nota1 + (nota2 * 2) + (nota3 * 3) + mediaExercicios) / 7
	switch {
	case mediaAproveitamento >= 9 && mediaAproveitamento <= 10:
		conceito = "A"
	case mediaAproveitamento >= 7.5 && mediaAproveitamento < 9:
		conceito = "B"
	case mediaAproveitamento >= 6 && mediaAproveitamento < 7.5:
		conceito = "C"
	case mediaAproveitamento >= 4 && mediaAproveitamento < 6:
		conceito = "D"
	default:
		conceito = "E"
	}
	fmt.Println("Dados do aluno")
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Nota 1: %.2f\n", nota1)
	fmt.Printf("Nota 2: %.2f\n", nota2)
	fmt.Printf("Nota 3: %.2f\n", nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaExercicios)
	fmt.Printf("Média de aproveitamento: %.2f\n", mediaAproveitamento)
	fmt.Printf("Conceito: %s\n", conceito)

	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}

