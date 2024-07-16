package mock

import (
	"fmt"
	"io"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3
const escrita =  "escrita"
const pausa = "pausa"

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SleeperPadrao struct {}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(saida, ultimaPalavra)
}


