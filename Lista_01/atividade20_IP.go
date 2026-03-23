package main
import "fmt"

func main() {
	var h, m, s int
	fmt.Scan(&h)
	fmt.Scan(&m)
	fmt.Scan(&s)

	total := h*3600 + m*60 + s
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", total)
}