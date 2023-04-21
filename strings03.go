// Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s)

	snovo := strings.ReplaceAll(s, "r", "l") //cebolinha
	fmt.Println(snovo)
}
