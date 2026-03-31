package main

import "fmt"

func main() {
    
    var n [5]float64
    var soma float64 = 0

    for i := 0; i < len(n); i++ {
        fmt.Printf("Digite a nota %d: ", i+1)
        fmt.Scan(&n[i])
        soma += n[i]
    }
        {
        fmt.Println("A soma total é de: ", soma)
    }
}
