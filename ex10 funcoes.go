package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarInteiros(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	sliceOrdenado := make([]int, len(slice))
	copy(sliceOrdenado, slice)

	sort.Ints(sliceOrdenado)
	return sliceOrdenado, nil
}

func main() {
	numeros := []int{9, 5, 2, 7, 1, 8}

	resultadoOrdenado, err := ordenarInteiros(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Valores ordenados:", resultadoOrdenado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := ordenarInteiros(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", resultadoVazio)
	}
}
