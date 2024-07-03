package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t testing.TB, resultado, esperado string)  {
		t.Helper() // informar que esse é um método auxiliar
		if resultado != esperado {
			t.Errorf("resultado %q, esperado %q", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Marie", "frances")
		esperado := "Bonjour, Marie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em inglês", func(t *testing.T) {
		resultado := Ola("John", "ingles")
		esperado := "Hello, John"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}