package services

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestJogoDaForca struct {
	suite.Suite
	jogo JogoDaForca
}

func TestJogoDaForcaPlay(t *testing.T){
suite.Run(t, new(TestJogoDaForca))
}

func (r *TestJogoDaForca) SetupTest() {
	r.jogo = NewJogoDaForca()
	// Arrange
	palavra:="macaco"

	// Action
	err := r.jogo.EntrarComPalavraEscondida(palavra)
	r.Require().NoError(err)
}

func (r *TestJogoDaForca) TestJogoDaForca_EntrarComPalavraEscondida_NaoDeveDarErro(){
	// Arrange
	palavra:="macaco"

	// Action
	escondida, err := r.jogo.GetPalavraEscondida()
	r.Require().NoError(err)

	//Assert
	r.Require().Equal(escondida, palavra)
}
func (r *TestJogoDaForca) TestJogoDaForca_AdicionarQuantidadeDeErros_MostrarQuantidadeCorreta(){
	// Arrange // Action
	r.jogo.AddQuantidadeErros()
	r.jogo.AddQuantidadeErros()
	r.jogo.AddQuantidadeErros()
	quantidadeErros,err := r.jogo.GetQuantidadeErros()
	r.Require().NoError(err)

	// Assert
	r.Require().Equal(quantidadeErros, 3)
}
func (r *TestJogoDaForca) TestJogoDaForca_EntrarComLetra_AoEntrarComLetraErrada_AdicionarQuantidadeErro(){
	// Arrange
	letraTentativa := "M"

	// Action
	r.jogo.EntrarComLetra(letraTentativa)
	quantidadeErros,err := r.jogo.GetQuantidadeErros()
	r.Require().NoError(err)

	// Assert
	r.Require().Equal(0, quantidadeErros)
}

func (r *TestJogoDaForca) TestJogoDaForca_EntrarComLetra_AoEntrarComLetraCerta_QuantidadeErroZero(){
	// Arrange
	letraTentativa := "P"

	// Action
	r.jogo.EntrarComLetra(letraTentativa)
	quantidadeErros,err := r.jogo.GetQuantidadeErros()
	r.Require().NoError(err)

	// Assert
	r.Require().Equal(1, quantidadeErros)
}

func (r *TestJogoDaForca) TestJogoDaForca_EscreverTentativa_MostrarEntradaDasLetras(){
	// Arrange
	listaTentativas:= "M-P"

	//Action
	r.jogo.EntrarComLetra("M")
	r.jogo.EntrarComLetra("P")

	stringListaTentativas := r.jogo.EscreverTentativa()

	// Assert
	r.Require().Equal(listaTentativas, stringListaTentativas)
}

func (r *TestJogoDaForca) TestJogoDaForca_DesenharForca(){
	// Arrange
	for i := 1; i < 13; i++ {
		r.jogo.AddQuantidadeErros()
	}

	//Action
	err:= r.jogo.DesenharForca()
	r.Require().NoError(err)
}

func (r *TestJogoDaForca) TestJogoDaForca_Integracao(){
	jogo := NewJogoDaForca()
	palavra:="macaco"
	err := jogo.EntrarComPalavraEscondida(palavra)
	if err != nil{
		panic("main: erro ao entrar com palavra escondida: %w")
	}
	err = jogo.EntrarComLetra("a")
	if err != nil{
		panic("main: erro ao entrar com a letra: %w")
	}
	jogo.DesenharForca()
}
