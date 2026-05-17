package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	var s float64
	for i := 1; i <= n; i++ {
		s += 1.0 / float64(i)
	}

	fmt.Printf("%.6f\n", s)
}