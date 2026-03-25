package main

import "fmt"

func main() {
    var alunos int
    var n, s, m float64

    fmt.Print("Digite a quantidade de alunos: ")
    fmt.Scan(&alunos)

    for i := 1; i <= alunos; i++ {
        fmt.Printf("Digite a nota do aluno %d: ", i)
        fmt.Scan(&n)
        s += n
    }

    if alunos > 0 {
        m = s / float64(alunos)
        fmt.Printf("\nA média da turma é: %.2f\n", m)
    } else {
        fmt.Println("Quantidade de alunos inválida.")
    }
}
