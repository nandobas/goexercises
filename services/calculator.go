package services

import (
	"fmt"
)

func ScanMultiplos(max int, other int, tipo string) int {

	result := 0

	for i := 1; i < max; i++ {

		result = SomaMultiplo(tipo, other, i, result)

		fmt.Println(result)
	}

	return result
}

func SomaMultiplo(tipo string, other int, i int, soma int) int {

	if tipo == "OR" {

		if i%3 == 0 {
			soma += i
		}

		if i%5 == 0 {
			soma += i
		}
	}
	if tipo == "AND" {

		if i%3 == 0 && i%5 == 0 {
			soma += i
		}

	}
	if tipo == "OR_AND_OTHER" {
		if (i%3 == 0 || i%5 == 0) && i%other == 0 {
			soma += i
		}
	}
	return soma
}

func Sum(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}
