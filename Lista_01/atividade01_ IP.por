programa {
  funcao inicio() {
    real v1, v2, v3, media

    escreva("Digite a primeira nota: ")
    leia(v1)

    enquanto (v1 < 0 ou v1 > 10) {
      escreva("Valor inválido, digite novamente: ")
      leia(v1)
    }

    escreva("Digite a segunda nota: ")
    leia(v2)

    enquanto (v2 < 0 ou v2 > 10) {
      escreva("Valor inválido, digite novamente: ")
      leia(v2)
    }

    escreva("Digite a terceira nota: ")
    leia(v3)

    enquanto (v3 < 0 ou v3 > 10) {
      escreva("Valor inválido, digite novamente: ")
      leia(v3)
    }

    media = (v1 + v2 + v3) / 3

    escreva("Sua média final é: ", media, "\n")

    se (media >= 6) {
      escreva("VOCÊ FOI APROVADO!\n")
    }
    senao {
      escreva("Tente de novo outra vez.\n")
    }
  }
}