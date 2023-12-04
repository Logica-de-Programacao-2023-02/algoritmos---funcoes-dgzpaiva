package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func stringsComMaiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	var resultado []string
	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {
			resultado = append(resultado, str)
		}
	}

	return strings.Join(resultado, " "), nil
}

func main() {
	strings := []string{"Gato", "Cachorro", "pássaro", "Elefante", "cobra"}

	resultado, err := stringsComMaiuscula(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com letra maiúscula:", resultado)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := stringsComMaiuscula(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", resultadoVazio)
	}
}
