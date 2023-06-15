// Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s)

	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")

	fmt.Println(s)

}
