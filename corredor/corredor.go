package corredor

import (
	"fmt"
	"net/http"
	"time"
)

var limiteDeDezSegundos = 10 * time.Second

func medirTempoDeResposta(URL string) time.Duration {
	inicio := time.Now()
	http.Get(URL)
	return time.Since(inicio)// pega o tempo inicial e retorna a diferença na forma de time.Duration.
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}

func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
	}
}

func Corredor(a, b string) (vencedor string, erro error) {
	return Configuravel(a, b, limiteDeDezSegundos)
}