package slices

import "testing"

func TestSoma(t *testing.T) {	
	
	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		esperado := Soma(numeros)
		resultado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})
}