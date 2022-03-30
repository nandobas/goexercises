package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertPalavraToNumero(t *testing.T) {
	require.Equal(t, ConvertPalavraToNumero("Mesa"), 123456)
}
