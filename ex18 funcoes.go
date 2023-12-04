package main

import (
	"errors"
	"fmt"
)

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func primosAteN(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("número negativo não é válido")
	}

	var primos []int
	for i := 2; i <= numero; i++ {
		if ehPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos, nil
}

func main() {
	numero := 20

	resultado, err := primosAteN(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos até", numero, ":", resultado)
	}

	numeroNegativo := -5
	resultadoNegativo, errNegativo := primosAteN(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Println("Resultado para número negativo:", resultadoNegativo)
	}
}
