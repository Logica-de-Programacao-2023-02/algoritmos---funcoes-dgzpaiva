package main

import (
	"errors"
	"fmt"
)

func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	pares := []int{}
	for _, num := range slice {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	return pares, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultadoPares, errPares := numerosPares(numeros)
	if errPares != nil {
		fmt.Println("Erro:", errPares)
	} else {
		fmt.Println("Números pares:", resultadoPares)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := numerosPares(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Números pares no slice vazio:", resultadoVazio)
	}
}
