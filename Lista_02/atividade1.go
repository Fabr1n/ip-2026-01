package main
import "fmt"

func main() {
    var numero int

    fmt.Print("Digite um número inteiro: ")
    fmt.Scan(&numero) 

    if numero%2 == 0 {
        fmt.Printf("O número %d é Par.\n", numero)
    } else {
        fmt.Printf("O número %d é Ímpar.\n", numero)
    }
}
