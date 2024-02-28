package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("Escolha a operação:")
	fmt.Println("1. Soma")
	fmt.Println("2. Subtração")
	fmt.Println("3. Divisão")
	fmt.Println("4. Multiplicação")

	var op int
	fmt.Scanln(&op)

	switch op {
	case 1:
		result := num1 + num2
		fmt.Printf("Resultado: %.2f\n", result)
	case 2:
		result := num1 - num2
		fmt.Printf("Resultado: %.2f\n", result)
	case 3:
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero")
		} else {
			result := num1 / num2
			fmt.Printf("Resultado: %.2f\n", result)
		}
	case 4:
		result := num1 * num2
		fmt.Printf("Resultado: %.2f\n", result)
	default:
		fmt.Println("Opção inválida")
	}
}
