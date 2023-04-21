// Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
// Informe ao usuário se é ou não.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("Escreva algo.")
	fmt.Scan(&s)

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("O valor %s não é um float \n.", s)
	} else {
		fmt.Printf("O valor float é %f\n.", i)
	}

}
