package main

import (
	"errors"
	"fmt"
)

func numerosComuns(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("um dos slices está vazio")
	}

	ocorrencias := make(map[int]bool)
	resultado := []int{}

	for _, num := range slice1 {
		ocorrencias[num] = true
	}

	for _, num := range slice2 {
		if ocorrencias[num] {
			resultado = append(resultado, num)
		}
	}

	return resultado, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	resultado, err := numerosComuns(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números presentes nos dois slices:", resultado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := numerosComuns(slice1, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o segundo slice vazio:", resultadoVazio)
	}
}
