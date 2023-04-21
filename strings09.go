// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.
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

	var letra string
	fmt.Println("Escolha uma letra indesejada na frase.")
	fmt.Scan(&letra)

	frasenova := strings.ReplaceAll(frase, letra, "")
	fmt.Println(frasenova)
}
