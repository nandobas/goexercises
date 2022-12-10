package services_test

import (
	"goexercises-app/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	b := services.PersonBuilder{}
	p := b.MyName("Nando").MyWork("dev").Build()
	require.Equal(t, "Nando", p.GetName())
	require.Equal(t, "dev", p.GetPosition())
}
