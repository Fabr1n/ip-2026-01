package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.14159265359

	fmt.Println("Volume de esferas com raio de 0 a 20 cm:")
	fmt.Println("Raio (cm) | Volume (cm³)")
	fmt.Println("------------------------")

	for raio := 0.0; raio <= 20.0; raio += 0.5 {
		volume := (4.0 / 3.0) * pi * math.Pow(raio, 3)
		fmt.Printf("%.1f      | %.2f\n", raio, volume)
	}
}
