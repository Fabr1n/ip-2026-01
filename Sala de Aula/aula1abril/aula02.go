package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	// Lê um número do tipo float
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	// Verifica se o número é positivo ou nulo (>= 0)
	if numero >= 0 {
		// Calcula a raiz quadrada
		raizQuadrada := math.Sqrt(numero)
		fmt.Printf("A raiz quadrada de %.2f é: %.4f\n", numero, raizQuadrada)
	} else {
		// Se for negativo, calcula o quadrado do número
		quadrado := numero * numero
		fmt.Printf("O número %.2f é negativo. Seu quadrado é: %.2f\n", numero, quadrado)
	}
}

/*
================== ANOTAÇÕES PARA APRENDIZADO ==================

1. IMPORTAÇÃO DE PACOTES:
   - import "pacote" importa um único pacote
   - import ( "pacote1"; "pacote2" ) importa múltiplos pacotes
   - "fmt": para formatação de entrada e saída
   - "math": pacote com funções matemáticas (Sqrt, Sin, Cos, etc)

2. FUNÇÃO math.Sqrt():
   - Calcula a raiz quadrada de um número
   - Sintaxe: math.Sqrt(numero)
   - Retorna um valor float64
   - Exemplo: math.Sqrt(16) retorna 4.0

3. ESTRUTURA CONDICIONAL - IF/ELSE:
   - Sintaxe: if condição { ... } else { ... }
   - Executa o bloco "if" se a condição for verdadeira
   - Executa o bloco "else" se a condição for falsa
   - Operadores de comparação: > (maior), < (menor), >= (maior ou igual), <= (menor ou igual), == (igual), != (diferente)

4. OPERAÇÕES MATEMÁTICAS:
   - Multiplicação (*): numero * numero calcula o quadrado
   - Adição (+), Subtração (-), Divisão (/)
   - Precedência: *, / executam antes de +, -

5. FLUXO DO PROGRAMA:
   1. Declara variável para armazenar o número
   2. Solicita e lê o número do usuário
   3. Verifica se o número é >= 0 (positivo ou nulo)
      - Se SIM: Calcula e exibe a raiz quadrada
      - Se NÃO (negativo): Calcula e exibe o quadrado do número

6. ESPECIFICADORES DE FORMATO UTILIZADOS:
   - %.2f: mostra 2 casas decimais
   - %.4f: mostra 4 casas decimais
   - Use a precisão apropriada conforme a necessidade

7. RAZÃO DE USAR RAIZ QUADRADA PARA POSITIVOS:
   - Raiz quadrada de número negativo não existe nos números reais
   - Por isso, para negativos, usamos o quadrado do número
   - Exemplo: raiz de -4 = indefinida, mas (-4)² = 16

8. EXEMPLO DE EXECUÇÃO:
   Entrada: 25
   Saída: A raiz quadrada de 25.00 é: 5.0000

   Entrada: -5
   Saída: O número -5.00 é negativo. Seu quadrado é: 25.00

9. BOAS PRÁTICAS:
   - Use nomes descritivos para variáveis (raizQuadrada, quadrado)
   - Importe apenas os pacotes que precisa
   - Verifique condições lógicas antes de operações matemáticas
   - Trate casos especiais (números negativos, zero, etc)

=================================================================
*/
