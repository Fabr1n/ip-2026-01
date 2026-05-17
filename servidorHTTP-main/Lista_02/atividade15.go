package main

import "fmt"

func mesPorExtenso(mes int) string {
	switch mes {
	case 1:
		return "janeiro"
	case 2:
		return "fevereiro"
	case 3:
		return "março"
	case 4:
		return "abril"
	case 5:
		return "maio"
	case 6:
		return "junho"
	case 7:
		return "julho"
	case 8:
		return "agosto"
	case 9:
		return "setembro"
	case 10:
		return "outubro"
	case 11:
		return "novembro"
	case 12:
		return "dezembro"
	default:
		return "mês inválido"
	}
}

func main() {
	var dia, mes, ano int
	fmt.Print("Digite o dia: ")
	fmt.Scanln(&dia)

	fmt.Print("Digite o mês: ")
	fmt.Scanln(&mes)

	fmt.Print("Digite o ano: ")
	fmt.Scanln(&ano)
	textoMes := mesPorExtenso(mes)
	fmt.Printf("Data formatada: %d de %s de %d\n", dia, textoMes, ano)
}

