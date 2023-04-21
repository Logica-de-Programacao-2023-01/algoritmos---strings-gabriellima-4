// Solicite ao usuário uma string e informe se ela contém pelo menos um número.
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

	fmt.Println(strings.ContainsAny(frase, "1,2,3,4,5,6,7,8,9,10"))

}
