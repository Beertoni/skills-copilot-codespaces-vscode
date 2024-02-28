package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite algo: ")
	input, _ := reader.ReadString('\n')

	// Remover espaços em branco
	input = strings.TrimSpace(input)

	// Contar caracteres
	count := len(input)

	fmt.Printf("O número de caracteres é: %d\n", count)
}
