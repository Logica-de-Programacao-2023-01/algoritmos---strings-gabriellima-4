//Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.
package main

import "fmt"

func main() {
	var s1 string
	var s2 string

	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s1)
	fmt.Print("Escreva uma palavra.")
	fmt.Scan(&s2)

	if s1 == s2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}

}
