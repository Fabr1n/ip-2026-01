package main

import "fmt"

// maiorDeTres recebe 3 números inteiros e retorna o maior deles.
func maiorDeTres(a, b, c int) int {
	maior := a

	// Compara o segundo valor com o maior atual.
	if b > maior {
		maior = b
	}

	// Compara o terceiro valor com o maior atual.
	if c > maior {
		maior = c
	}

	return maior
}

func main() {
	var num1, num2, num3 int

	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número inteiro: ")
	fmt.Scanln(&num3)

	maior := maiorDeTres(num1, num2, num3)

	fmt.Printf("O maior número é: %d\n", maior)
}

/*
================== ANOTAÇÕES PARA APRENDIZADO ==================

1. FUNÇÃO EM GO:
	- Sintaxe: func nome(parametro tipo, parametro tipo) tipo_retorno { ... }
	- A função maiorDeTres() recebe 3 inteiros e retorna 1 inteiro
	- Ela compara os valores e guarda o maior na variável maior

2. TIPOS DE DADOS:
	- int: número inteiro
	- Ideal para contagens, índices e valores sem casas decimais

3. DECLARAÇÃO DE VARIÁVEIS:
	- var num1, num2, num3 int declara 3 variáveis do tipo inteiro
	- var maior := maiorDeTres(...) cria uma variável já com valor inicial

4. ENTRADA DE DADOS (INPUT):
	- fmt.Scanln(&variavel) lê um valor digitado pelo usuário
	- O & passa o endereço da variável para que o valor seja alterado

5. SAÍDA DE DADOS (OUTPUT):
	- fmt.Print() mostra a mensagem sem quebra de linha
	- fmt.Printf() permite formatar a saída com especificadores
	- %d é usado para números inteiros

6. ESTRUTURA CONDICIONAL - IF:
	- Sintaxe: if condição { ... }
	- O bloco é executado somente se a condição for verdadeira
	- Neste exercício, usamos if para comparar qual número é maior

7. LÓGICA DO ALGORITMO:
	1. Assume que o primeiro número é o maior
	2. Compara o segundo número com o maior atual
	3. Se o segundo for maior, ele substitui o valor guardado
	4. Compara o terceiro número com o maior atual
	5. Ao final, a variável maior contém o maior dos 3 valores

8. EXEMPLO DE EXECUÇÃO:
	Entrada: 10, 25, 7
	Saída: O maior número é: 25

9. BOAS PRÁTICAS:
	- Use nomes descritivos para funções e variáveis
	- Separe a lógica principal em uma função para facilitar a leitura
	- Faça comparações simples e objetivas

=================================================================
*/
