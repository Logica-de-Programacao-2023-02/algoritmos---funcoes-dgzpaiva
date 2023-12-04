package main

import (
	"fmt"
	"sort"
)

func segundoMaior(slice []int) int {
	if len(slice) < 2 {
		return -1 
	}

	sort.Ints(slice)
	maior := slice[len(slice)-1]
	for i := len(slice) - 2; i >= 0; i-- {
		if slice[i] < maior {
			return slice[i]
		}
	}
	return -1 
}

func main() {
	numeros := []int{10, 5, 8, 20, 15}
	segundo := segundoMaior(numeros)
	fmt.Println("O segundo maior valor é:", segundo)
}
