package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordemAlfabetica(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("o slice está vazio")
	}

	sort.Strings(slice)
	novaString := strings.Join(slice, " ")
	return novaString, nil
}

func main() {
	slice := []string{"maçã", "banana", "laranja", "abacaxi"}

	stringOrdenada, err := ordemAlfabetica(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings em ordem alfabética:", stringOrdenada)
	}

	sliceVazio := []string{}
	stringVazia, errVazio := ordemAlfabetica(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado para o slice vazio:", stringVazia)
	}
}
