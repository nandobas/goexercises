package main

import (
	"fmt"
	"services/services"
)

func exibePainel(j services.JogoDaForca){

	j.DesenharForca()
	fmt.Printf("Entativas: %s",	j.EscreverTentativa())
	qtdErros, _ := j.GetQuantidadeErros()
	fmt.Printf("\nErros: %v \n",	qtdErros)
	if qtdErros>=7 {
		fmt.Println("ENFORCADO!")
	}
}

func main() {
	jogo := services.NewJogoDaForca()
	palavra := "macaco"
	jogo.EntrarComPalavraEscondida(palavra)
	jogo.EntrarComLetra("a")
	jogo.EntrarComLetra("e")
	jogo.EntrarComLetra("i")
	jogo.EntrarComLetra("g")
	jogo.EntrarComLetra("p")
	jogo.EntrarComLetra("r")
	jogo.EntrarComLetra("c")
	jogo.EntrarComLetra("m")
	jogo.EntrarComLetra("u")
	jogo.EntrarComLetra("v")
	jogo.EntrarComLetra("o")
	exibePainel(jogo)

}
