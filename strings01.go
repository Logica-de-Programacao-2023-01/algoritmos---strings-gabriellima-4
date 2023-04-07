// Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas.
// Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Digite uma palavra.")
	fmt.Scan(&s)

	fmt.Println(s)

	fmt.Println(strings.ToUpper(s))
}
