package main

import "fmt"

func main() {
	var numAlunos int

	fmt.Print("Digite a quantidade de alunos: ")
	fmt.Scanln(&numAlunos)

	var aprovados, exame, reprovados int
	somaNotas := 0.0

	for i := 1; i <= numAlunos; i++ {
		var nota1, nota2 float64

		fmt.Printf("\nAluno %d:\n", i)
		fmt.Print("Digite a primeira nota: ")
		fmt.Scanln(&nota1)

		fmt.Print("Digite a segunda nota: ")
		fmt.Scanln(&nota2)

		media := (nota1 + nota2) / 2
		somaNotas += media

		fmt.Printf("Média: %.2f - ", media)

		if media < 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media >= 3 && media < 7 {
			fmt.Println("Exame")
			exame++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}

	mediaClasse := somaNotas / float64(numAlunos)

	fmt.Printf("\n--- RESUMO ---\n")
	fmt.Printf("c) Total de alunos aprovados: %d\n", aprovados)
	fmt.Printf("d) Total de alunos de exame: %d\n", exame)
	fmt.Printf("e) Total de alunos reprovados: %d\n", reprovados)
	fmt.Printf("f) Média da classe: %.2f\n", mediaClasse)
}
