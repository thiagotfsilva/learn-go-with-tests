package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes map[string]int
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	req , _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisição é inválido, obtive '%s' esperava '%s'", recebido, esperado)
}
}

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
	}
	servidor := &ServidorJogador{&armazenamento}
	t.Run("retorna resultado de Maria", func(t *testing.T) {
		req := novaRequisicaoObterPontuacao("Maria")
		res := httptest.NewRecorder()

		servidor.ServeHTTP(res, req)

		verificarCorpoRequisicao(t, res.Body.String(), "20")
	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		req := novaRequisicaoObterPontuacao("Pedro")
		res := httptest.NewRecorder()

		servidor.ServeHTTP(res, req)

		verificarCorpoRequisicao(t, res.Body.String(), "10")
	})
}