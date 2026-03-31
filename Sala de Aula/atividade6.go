package main

import "fmt"

func main() {

    var nota [5]float64
    
    for i := 0; i < len(nota); i++ {
        fmt.Printf("Digite a nota %d: ", i+1)
        fmt.Scan(&nota[i])
    }
    
    fmt.Println("\nNotas em ordem reversa:")
    for i := len(nota) - 1; i >= 0; i-- {
        fmt.Println(nota[i])
    }
}
