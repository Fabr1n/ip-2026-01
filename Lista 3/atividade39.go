package main

import "fmt"

func main() {
	var totalBois int = 90

	fmt.Println("Digite o peso de cada boi:")

	var maiorPeso, menorPeso int
	var maiorgordoID, maiormagroID int

	for i := 1; i <= totalBois; i++ {
		fmt.Printf("Boi %d - Peso: ", i)
		var peso int
		fmt.Scanln(&peso)

		if i == 1 {
			maiorPeso = peso
			menorPeso = peso
			maiorgordoID = i
			maiormagroID = i
		} else {
			if peso > maiorPeso {
				maiorPeso = peso
				maiorgordoID = i
			}
			if peso < menorPeso {
				menorPeso = peso
				maiormagroID = i
			}
		}
	}

	fmt.Printf("\nBoi mais gordo - ID: %d, Peso: %d kg\n", maiorgordoID, maiorPeso)
	fmt.Printf("Boi mais magro - ID: %d, Peso: %d kg\n", maiormagroID, menorPeso)
}
