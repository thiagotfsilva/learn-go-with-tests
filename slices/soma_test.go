package slices

import (
	"reflect"
	"testing"
)

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

func TestSomaTudo(t *testing.T) {
	t.Run("soma os itens de variso slices e devolve um novo slice", func(t *testing.T) {
		resultado := SomaTudo([]int{1, 2}, []int{0, 9})
		esperado := []int{3, 9}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v esperado %v", resultado, esperado)
		}
	})
}

func TestSomaTodoResto(t *testing.T) {
	t.Run("soma o total de todos os finais de cada slice", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v esperado %v", resultado, esperado)
		}
	})

	t.Run("Soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v esperado %v", resultado, esperado)
		}
	})
}