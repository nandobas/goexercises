package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSouNumeroFeliz7(t *testing.T) {

	var lista_dos_que_ja_passaram []int64
	lista_dos_que_ja_passaram = append(lista_dos_que_ja_passaram, 0)
	require.Equal(t, SouNumeroFeliz(7, lista_dos_que_ja_passaram), true)

}

func TestSouNumeroFeliz40(t *testing.T) {

	var lista_dos_que_ja_passaram []int64
	lista_dos_que_ja_passaram = append(lista_dos_que_ja_passaram, 0)
	require.Equal(t, SouNumeroFeliz(40, lista_dos_que_ja_passaram), false)

}
