package main

import "fmt"

func main() {
    var x float64

    fmt.Print("Digite o número: ")
    fmt.Scan(&x)

    if x > 0 {
        fmt.Println("O número", x, "é positivo.")
    } else if x < 0 {
        fmt.Println("O número", x, "é negativo.")
    } else {
        fmt.Println("O número", x, "é neutro.")
    }
}