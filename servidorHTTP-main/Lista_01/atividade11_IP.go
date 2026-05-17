package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}