package ponteiroserros

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin)  {
		t.Helper()
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	confirmaErro := func(t *testing.T, resultado error, esperado error)  {
		t.Helper()

		if resultado == nil {
			t.Error("esperava um erro, mas nenhum ocorreu.")
		}

		if resultado != esperado {
			t.Errorf("erro resultado %s, erro esperado %s", resultado, esperado)
		}
	}

	confirmaErroInexistente := func(t*testing.T, resultado error)  {
		t.Helper()
		if resultado != nil {
			t.Fatal("erro inesperado recebido")
		}
	}

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))
		
		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}