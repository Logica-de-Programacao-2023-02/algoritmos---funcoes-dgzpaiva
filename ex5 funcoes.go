package main

import (
	"fmt"
)

func encontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	numeros := []int{3, 7, 11, 5, 9, 13}
	valor := 5
	posicao := encontrarPosicao(numeros, valor)
	fmt.Printf("O primeiro %d está na posição: %d\n", valor, posicao)
}
