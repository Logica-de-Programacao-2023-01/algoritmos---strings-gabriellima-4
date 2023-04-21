// Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var s2 string

	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s1)
	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s2)

	if strings.Contains(s1, s2) {
		fmt.Println("A string s2 é substring da primeira.")
	} else {
		fmt.Println("A string s1 não contém a string s2.")
	}
}
