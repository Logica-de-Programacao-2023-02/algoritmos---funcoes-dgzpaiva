package main

import (
	"errors"
	"fmt"
	"math"
)

func ehPrimo(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("número negativo não é válido")
	}

	if numero <= 1 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero))) + 1
	for i := 2; i < limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	numero := 17

	primo, err := ehPrimo(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else if primo {
		fmt.Println(numero, "é um número primo")
	} else {
		fmt.Println(numero, "não é um número primo")
	}

	numeroNegativo := -5
	primoNegativo, errNegativo := ehPrimo(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Println("Resultado para número negativo:", primoNegativo)
	}
}
