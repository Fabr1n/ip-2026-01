programa {
  funcao inicio() {

    inteiro casos
    escreva("qual o n do jogo?: ")
    leia(casos)

    real pessoas
    escreva("Quantas pessoas participaram do jogo?: ")
    leia(pessoas)

    real perPopular
    real perGeral
    real perArquibancada
    real perCadeiras

    escreva("digite o percentual de ingressos populares vendidos: ")
    leia(perPopular)

    escreva("digite o percentual de ingressos gerais vendidos: ")
    leia(perGeral)

    escreva("digite o percentual de ingressos arquibancadas vendidos: ")
    leia(perArquibancada)

    escreva("digite o percentual de ingressos cadeiras vendidos: ")
    leia(perCadeiras)

    real vpopular
    real vgeral
    real varquibancadas
    real vcadeiras

    vpopular = 1
    vgeral = 5
    varquibancadas = 10
    vcadeiras = 20

    real rpopular
    real rgeral
    real rarquibancadas
    real rcadeiras

    rpopular = vpopular * (pessoas * (perPopular / 100))
    rgeral = vgeral * (pessoas * (perGeral / 100))
    rarquibancadas = varquibancadas * (pessoas * (perArquibancada / 100))
    rcadeiras = vcadeiras * (pessoas * (perCadeiras / 100))

    real renda
    renda = rpopular + rgeral + rarquibancadas + rcadeiras
    escreva("A RENDA TOTAL DO JOGO N. ", casos, " E = R$", renda)
    }

  }

