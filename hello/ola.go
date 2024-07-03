package main

import "fmt"

const prefixoOla = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hello, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	
	return  prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = prefixoOlaFrances
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "ingles":
		prefixo = prefixoOlaIngles
	default: 
		prefixo = prefixoOla
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}