package main

import "fmt"

func main() {
    
    var n [3]float64
    var soma float64 = 0

    for i := 0; i < len(n); i++ {
        fmt.Printf("Digite a nota %d: ", i+1)
        fmt.Scan(&n[i])
        soma += n[i]
    }

    media := soma / float64(len(n))
    fmt.Printf("\nMédia Geral: %.2f\n", media)

    fmt.Println("Notas acima da média:")
    for _, nota := range n {
        if nota > media {
            fmt.Printf("- %.2f\n", nota)
        }
    }
}
