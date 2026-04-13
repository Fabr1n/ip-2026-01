package main

import "fmt"

func main() {
	var totalPessoas, maiores50, entreCincode20 int
	var menorQue40 int
	var somaAlturasCincoVinte, somaAlturas float64

	for {
		var idade int
		var altura, peso float64

		fmt.Print("Digite a idade: ")
		fmt.Scanln(&idade)

		fmt.Print("Digite a altura (m): ")
		fmt.Scanln(&altura)

		fmt.Print("Digite o peso (kg): ")
		fmt.Scanln(&peso)

		totalPessoas++

		if idade > 50 {
			maiores50++
		}

		if idade >= 10 && idade <= 20 {
			somaAlturasCincoVinte += altura
			entreCincode20++
		}

		if peso < 40 {
			menorQue40++
		}

		fmt.Print("Deseja continuar? (1-Sim, outro-Não): ")
		var continua int
		fmt.Scanln(&continua)

		if continua != 1 {
			break
		}
	}

	fmt.Printf("\nPessoas com idade superior a 50 anos: %d\n", maiores50)

	if entreCincode20 > 0 {
		mediaAltura := somaAlturasCincoVinte / float64(entreCincode20)
		fmt.Printf("Média de alturas (10-20 anos): %.2f m\n", mediaAltura)
	} else {
		fmt.Println("Nenhuma pessoa entre 10 e 20 anos")
	}

	percentual := (float64(menorQue40) / float64(totalPessoas)) * 100
	fmt.Printf("Porcentagem de pessoas com peso < 40kg: %.2f%%\n", percentual)
}
