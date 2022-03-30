package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	/*if Sum(2, 2) != 4 {
		t.Error("2 + 2 must be 4")
	}*/
	require.Equal(t, Sum(2, 2), 4)

}
func TestMultiply(t *testing.T) {
	/*if Multiply(2, 2) != 4 {
		t.Error("2 * 2 must be 4")
	}*/

	require.Equal(t, Multiply(2, 2), 4)
}

func TestSomaMultiplosAbaixo10(t *testing.T) {

	require.Equal(t, ScanMultiplos(10, 0, "OR"), 23)

}

func TestSomaMultiplosAbaixo1000TipoOr(t *testing.T) {

	require.Equal(t, ScanMultiplos(1000, 0, "OR"), 266333)

}

func TestSomaMultiplosAbaixo1000TipoAnd(t *testing.T) {

	require.Equal(t, ScanMultiplos(1000, 0, "AND"), 33165)

}

func TestSomaMultiplosAbaixo1000TipoOrAndOther(t *testing.T) {

	require.Equal(t, ScanMultiplos(1000, 7, "OR_AND_OTHER"), 33173)

}
