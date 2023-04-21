// Escreva um programa que receba uma string e remova todos os espaços em branco.
// Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("Hoje é sexta feira.", " ", ""))
}
