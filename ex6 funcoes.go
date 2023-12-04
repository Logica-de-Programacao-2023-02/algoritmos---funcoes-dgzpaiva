package main

import (
	"errors"
	"fmt"
	"strings"
)

func concatenarComVirgulas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}
	return strings.Join(slice, ", "), nil
}

func main() {
	palavras := []string{"Olá", "isso", "é", "um", "exemplo"}
	resultado, err := concatenarComVirgulas(palavras)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Concatenação:", resultado)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := concatenarComVirgulas(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Concatenação:", resultadoVazio)
	}
}
