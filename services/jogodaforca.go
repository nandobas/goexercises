package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// INTERFACES
type JogoDaForca interface{
	EntrarComPalavraEscondida(palavra string) error
	GetPalavraEscondida() (string, error)
	EntrarComLetra(letra string) error
	AddQuantidadeErros()
	GetQuantidadeErros()(int, error)
	EscreverTentativa()string
	DesenharForca() error
}

type tentativa interface {
	EntraTentativa(letra string) error
	ListaTentativas()[]string
}

// IMPLEMENTS
type MinhasTentativas struct {
	listaTentativas []string
}

func (m *MinhasTentativas) EntraTentativa(letra string) error {

	chars := []rune(letra)
	listaTentativas := strings.Join(m.listaTentativas, "")
	index := strings.IndexRune(listaTentativas, chars[0])
	if index <= 0 {
		m.listaTentativas= append(m.listaTentativas, letra)
	}
	return nil
}

func (m *MinhasTentativas) ListaTentativas() []string{
	return m.listaTentativas
}

func newTentativas() tentativa{
	return &MinhasTentativas{}
}

type jogoDaForca struct {
	contadorErros int
	palavraEscondida []string
	tentativas tentativa
}

func (j *jogoDaForca) AddQuantidadeErros() {
	j.contadorErros++
}

func (j *jogoDaForca) GetQuantidadeErros() (int, error) {
	if j.contadorErros<0 {
		return 0, fmt.Errorf("Por algum motivo a quantidade de erros esta abaixo de zero %d", j.contadorErros)
	}
	return j.contadorErros, nil

}

func NewJogoDaForca() JogoDaForca{
	return &jogoDaForca{contadorErros: 0, palavraEscondida: make([]string, 1), tentativas: newTentativas()}
}

func (j *jogoDaForca) EntrarComPalavraEscondida(palavra string) error{
	j.palavraEscondida = strings.Split(palavra, "")
	return nil
}

func (j *jogoDaForca) GetPalavraEscondida() (string, error) {
	if len(j.palavraEscondida) <=0 {
		return "",fmt.Errorf("Palavra escondida não foi cadastrada")
	}
	return strings.Join(j.palavraEscondida, ""), nil
}

func (j *jogoDaForca) EntrarComLetra(letra string) error {

	palavra, err := j.GetPalavraEscondida()
	if err!=nil{
		return fmt.Errorf("Palavra não localizada: %w", err)
	}

	chars := []rune(letra)

	index := strings.IndexRune(palavra, chars[0])
	if index < 0{
		j.AddQuantidadeErros()
	}

	j.tentativas.EntraTentativa(letra)
	return nil
}

func (j *jogoDaForca) EscreverTentativa() string {
	return strings.Join(j.tentativas.ListaTentativas(), "-")
}

func (j *jogoDaForca) DesenharForca() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	var desenhoForca [14]string
	desenhoForca[0]  = `     |      `
	desenhoForca[1]  = `     |      `
	desenhoForca[2]  = `    _|__    `
	desenhoForca[3]  = `   /    \   `
	desenhoForca[4]  = `   \    /   `
	desenhoForca[5]  = ` /--------\ `
	desenhoForca[6]  = ` |  ----   |`
	desenhoForca[7]  = ` |  ----   |`
	desenhoForca[8]  = `    -  -    `
	desenhoForca[9]  = `   /    \   `
	desenhoForca[10] = `  /      \  `
	desenhoForca[11] = `  |      |  `
	desenhoForca[12] = `__|      |__`
	desenhoForca[13] = `            `

	for i := 0; i < j.contadorErros*2; i=i+2 {

		fmt.Println(desenhoForca[i]+"\n"+desenhoForca[i+1])
	}
	return nil
}

func entrarComPalavraSecreta(args ...string) error {
	fmt.Println("Escolha uma palavra secreta")
	return nil
}

func entrarComLetra(args ...string) error {
	fmt.Println("Digite uma letra")
	return nil
}