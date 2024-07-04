package iteracao

func Repertir(caractere string, iteracoes int) string {
	var repeticoes string
	for i := 0; i < iteracoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
