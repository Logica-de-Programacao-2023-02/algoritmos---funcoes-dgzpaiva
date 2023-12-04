package main

import (
	"errors"
	"fmt"
	"strconv"
)

func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("número negativo não é válido")
	}

	digitos := strconv.Itoa(numero)
	soma := 0
	for _, digito := range digitos {
		d, _ := strconv.Atoi(string(digito))
		soma += d
	}
	return soma, nil
}

func main() {
	numero := 12345

	resultado, err := somaDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos:", resultado)
	}

	numeroNegativo := -54321
	resultadoNegativo, errNegativo := somaDigitos(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Println("Resultado para número negativo:", resultadoNegativo)
	}
}
