package main

import (
	"errors"
	"fmt"
)

func stringsMaisDe5Caracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	var resultado []string
	for _, str := range slice {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	strings := []string{"gato", "cachorro", "passarinho", "elefante", "cobra"}

	resultado, err := stringsMaisDe5Caracteres(strings)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", resultado)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := stringsMaisDe5Caracteres(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", resultadoVazio)
	}
}
