// Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println(strings.Count(frase, " ") + 1)

}
