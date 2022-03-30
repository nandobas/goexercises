package services

import (
	"strconv"
)

func SouNumeroFeliz(numero int64, lista_dos_que_ja_passaram []int64) bool {

	var soma int64
	soma = 0

	numeros := convertToArray(numero)
	for i := 0; i < len(numeros); i++ {
		char := (numeros[i])
		soma += (char * char)
	}

	if soma == 1 {
		return true
	}

	var b_search = verificaListaJaPassaram(soma, lista_dos_que_ja_passaram)
	if b_search {
		return false
	}

	lista_dos_que_ja_passaram = append(lista_dos_que_ja_passaram, soma)
	return SouNumeroFeliz(soma, lista_dos_que_ja_passaram)
}

func verificaListaJaPassaram(find int64, lista_dos_que_ja_passaram []int64) bool {
	for i := 0; i < len(lista_dos_que_ja_passaram); i++ {
		char := (lista_dos_que_ja_passaram[i])
		if char == find {
			return true
		}
	}
	return false
}

func convertToArray(numero int64) [4]int64 {

	var result [4]int64

	str := strconv.FormatInt(numero, 10)

	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		x, err := strconv.ParseInt(char, 10, 64)
		if err == nil {
			result[i] = x
		}
	}
	return result
}
