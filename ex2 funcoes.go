package main

import (
	"fmt"
	"strings"
)

func contarVogais(s string) int {
	vogais := "aeiouAEIOU"
	contagem := 0
	for _, char := range s {
		if strings.ContainsRune(vogais, char) {
			contagem++
		}
	}
	return contagem
}

func main() {
	texto := "Olá, isso é um exemplo de contagem de vogais!"
	qtdVogais := contarVogais(texto)
	fmt.Printf("O texto tem %d vogais.\n", qtdVogais)
}
