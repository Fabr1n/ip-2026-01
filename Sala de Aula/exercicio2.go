package main
import "fmt"

func main() {
    var nota1, peso1, nota2, peso2 int
  fmt.Println("Digite o valor da Nota 1: ")
  fmt.Scan(&nota1)
  fmt.Println("Digite o valor do Peso 1: ")
  fmt.Scan(&peso1)
  fmt.Println("Digite o valor da Nota 2: ")
  fmt.Scan(&nota2)
  fmt.Println("Digite o valor da Pes 1: ")
  fmt.Scan(&peso2)

fmt.Println("A média ponderada é de:", ((nota1) * (peso1) + (nota2) * (peso2) / ((peso1) + (peso2))))
}