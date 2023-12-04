package main

import (
	"errors"
	"fmt"
	"strings"
)

func extrairPalavras(texto string) ([]string, error) {
	if texto == "" {
		return nil, errors.New("a string está vazia")
	}

	palavras := strings.Fields(texto)
	return palavras, nil
}

func main() {
	frase := "Esta é uma frase de exemplo"
	resultado, err := extrairPalavras(frase)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", resultado)
	}

	stringVazia := ""
	resultadoVazio, errVazio := extrairPalavras(stringVazia)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para string vazia:", resultadoVazio)
	}
}
