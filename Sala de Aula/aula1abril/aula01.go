package main

import "fmt"

// Função que recebe 2 valores e retorna a soma
func soma(a float64, b float64) float64 {
	return a + b
}

func main() {
	var num1, num2 float64

	// Lê o primeiro número
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	// Lê o segundo número
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	// Calcula a soma através da função
	resultado := soma(num1, num2)

	// Exibe o resultado
	fmt.Printf("A soma de %.2f + %.2f = %.2f\n", num1, num2, resultado)
}

/*
================== ANOTAÇÕES PARA APRENDIZADO ==================

1. FUNÇÕES EM GO:
   - Sintaxe: func nome(parametro tipo, parametro tipo) tipo_retorno { ... }
   - A função soma() recebe 2 parâmetros (a e b) do tipo float64
   - Retorna um valor do tipo float64
   - Múltiplos parâmetros do mesmo tipo podem ser escritos: (a, b float64)

2. TIPOS DE DADOS:
   - float64: número decimal com ponto flutuante (precisão de 64 bits)
   - Ideal para operações matemáticas com decimais
   - Alternativa: int para números inteiros, float32 para menor precisão

3. DECLARAÇÃO DE VARIÁVEIS:
   - Sintaxe: var nome tipo
   - var num1, num2 float64 declara 2 variáveis do mesmo tipo
   - Variáveis em Go são inicializadas com valor padrão (0 para números)

4. ENTRADA DE DADOS (INPUT):
   - fmt.Scanln(&variable) lê uma entrada do usuário
   - O & (ampersand) passa o endereço da variável em memória
   - Permite que a função modifique o valor da variável

5. SAÍDA DE DADOS (OUTPUT):
   - fmt.Print() exibe sem quebra de linha
   - fmt.Println() exibe com quebra de linha
   - fmt.Printf() formata a saída com especificadores (%.2f = 2 casas decimais)

6. ESPECIFICADORES DE FORMATO:
   - %.2f: mostra o número decimal com 2 casas após a vírgula
   - %d: inteiro
   - %s: string
   - %v: valor genérico

7. FLUXO DO PROGRAMA:
   1. Declara 2 variáveis para armazenar os números
   2. Solicita e lê o primeiro número do usuário
   3. Solicita e lê o segundo número do usuário
   4. Chama a função soma() passando os 2 números
   5. Armazena o resultado retornado pela função
   6. Exibe o resultado formatado

8. BOAS PRÁTICAS:
   - Use nomes descritivos para variáveis e funções
   - Adicione comentários para explicar partes complexas
   - Utilize tipos apropriados para cada situação
   - Mantenha as funções simples e com apenas uma responsabilidade

=================================================================
*/
