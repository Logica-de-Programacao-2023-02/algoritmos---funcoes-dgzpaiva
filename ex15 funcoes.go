package main

import (
	"errors"
	"fmt"
)

func numeroPresente(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("o slice está vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	numero := 5
	slice := []int{2, 4, 6, 8, 10}

	presente, err := numeroPresente(numero, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("O número está presente no slice:", presente)
	}

	sliceVazio := []int{}
	presenteVazio, errVazio := numeroPresente(numero, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", presenteVazio)
	}
}
