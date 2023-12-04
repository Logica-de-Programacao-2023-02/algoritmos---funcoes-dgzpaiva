package main

import (
	"fmt"
	"strings"
)

func concatenarStrings(slice []string) string {
	return strings.Join(slice, "")
}

func main() {
	palavras := []string{"Olá", ", ", "isso", " ", "é", " ", "um", " ", "exemplo", "!"}
	resultado := concatenarStrings(palavras)
	fmt.Println("Concatenação:", resultado)
}
