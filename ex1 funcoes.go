package main

import (
	"fmt"
)

func calcularMedia(slice []int) float64 {
	total := 0
	for _, valor := range slice {
		total += valor
	}
	return float64(total) / float64(len(slice))
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	media := calcularMedia(numeros)
	fmt.Printf("A média é: %.2f\n", media)
}
