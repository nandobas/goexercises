package services

import (
	"fmt"
	"testing"
)

func TestBuildCar (t *testing.T){
	builder := CarBuilder{}
	car := builder.Color("Black").FactoryAndModel("VW", "Gol").Year(2020).Build()
	fmt.Println(*car)
}
