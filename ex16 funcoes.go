package main

import (
	"errors"
	"fmt"
	"strings"
)

func substituirVogaisPorAsterisco(texto string) (string, error) {
	if texto == "" {
		return "", errors.New("a string está vazia")
	}

	vogais := "aeiouAEIOU"
	var resultado strings.Builder

	for _, char := range texto {
		if strings.ContainsRune(vogais, char) {
			resultado.WriteRune('*')
		} else {
			resultado.WriteRune(char)
		}
	}

	return resultado.String(), nil
}

func main() {
	texto := "Olá, isso é um exemplo de substituição de vogais!"
	
	novoTexto, err := substituirVogaisPorAsterisco(texto)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Novo texto:", novoTexto)
	}

	textoVazio := ""
	novoTextoVazio, errVazio := substituirVogaisPorAsterisco(textoVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para string vazia:", novoTextoVazio)
	}
}
