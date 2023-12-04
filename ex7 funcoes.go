package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("o slice está vazio")
	}

	resultados := make([]int, len(slice))
	for i, valor := range slice {
		resultados[i] = funcao(valor)
	}
	return resultados, nil
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	duplicar := func(x int) int {
		return x * 2
	}

	resultadoDuplicado, errDuplicado := aplicarFuncao(numeros, duplicar)
	if errDuplicado != nil {
		fmt.Println("Erro:", errDuplicado)
	} else {
		fmt.Println("Resultado após aplicar a função de duplicação:", resultadoDuplicado)
	}

	aoQuadrado := func(x int) int {
		return x * x
	}

	resultadoAoQuadrado, errAoQuadrado := aplicarFuncao(numeros, aoQuadrado)
	if errAoQuadrado != nil {
		fmt.Println("Erro:", errAoQuadrado)
	} else {
		fmt.Println("Resultado após aplicar a função de elevar ao quadrado:", resultadoAoQuadrado)
	}
	sliceVazio := []int{}
	resultadoVazio, errVazio := aplicarFuncao(sliceVazio, duplicar)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", resultadoVazio)
	}
}
