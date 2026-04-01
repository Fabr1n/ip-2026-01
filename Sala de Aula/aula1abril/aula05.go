package main

import "fmt"

// Fatorial recebe um número inteiro e retorna o seu fatorial.
func Fatorial(n int) int {
	resultado := 1

	for i := 2; i <= n; i++ {
		resultado *= i
	}

	return resultado
}

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	fatorial := Fatorial(numero)

	fmt.Printf("O fatorial de %d é: %d\n", numero, fatorial)
}

/*
================== ANOTAÇÕES PARA APRENDIZADO ==================

1. FUNÇÃO EM GO:
   - Sintaxe: func nome(parametro tipo) tipo_retorno { ... }
   - A função Fatorial() recebe 1 número inteiro e retorna 1 inteiro
   - Ela calcula o produto dos números de 1 até n

2. TIPOS DE DADOS:
   - int: número inteiro
   - Ideal para contagens, índices e cálculos sem decimais

3. DECLARAÇÃO DE VARIÁVEIS:
   - var numero int declara uma variável inteira
   - var fatorial := Fatorial(numero) guarda o resultado retornado

4. ENTRADA DE DADOS (INPUT):
   - fmt.Scanln(&variavel) lê o valor digitado pelo usuário
   - O & permite que a função de leitura preencha a variável

5. SAÍDA DE DADOS (OUTPUT):
   - fmt.Print() mostra a mensagem sem quebra de linha
   - fmt.Printf() mostra o resultado com formatação
   - %d é usado para valores inteiros

6. LAÇO DE REPETIÇÃO FOR:
   - Sintaxe: for inicialização; condição; atualização { ... }
   - Neste exercício, o laço percorre de 2 até n
   - Em cada passo, multiplica o resultado pelo valor atual de i

7. LÓGICA DO FATORIAL:
   - 0! = 1 e 1! = 1
   - Para n > 1, multiplica-se: 1 * 2 * 3 * ... * n
   - Exemplo: 5! = 1 * 2 * 3 * 4 * 5 = 120

8. FLUXO DO PROGRAMA:
   1. Lê um número inteiro
   2. Chama a função Fatorial()
   3. Recebe o valor calculado
   4. Exibe o resultado formatado

9. BOAS PRÁTICAS:
   - Inicialize o acumulador com 1 no cálculo do fatorial
   - Use um laço simples e claro para multiplicar os valores
   - Separe a lógica do cálculo da entrada e da saída

=================================================================
*/
