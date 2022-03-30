package services_test

import (
	"github.com/stretchr/testify/require"
	"services/services"
	"testing"
)

func TestBuild(t *testing.T) {
	b := services.PersonBuilder{}
	p := b.MyName("Nando").MyWork("dev").Build()
	require.Equal(t, "Nando", p.GetName())
	require.Equal(t, "dev", p.GetPosition())
}

