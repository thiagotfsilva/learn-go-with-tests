package estrutura

import "testing"

func TestArea(t *testing.T) {
	/* verificaArea := func (t *testing.T, forma Forma, esperado float64)  {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		verificaArea(t, retangulo, 72)
	})

	t.Run("círculos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	}) */
	t.Run("círculos", func(t *testing.T) {
		testesArea := []struct {
			nome string
			forma Forma
			areaEsperada  float64
		}{
			{forma: Retangulo{Lagura: 12, Altura:  6}, areaEsperada: 72.0},
			{forma: Circulo{Raio: 10}, areaEsperada: 314.1592653589793},
			{forma: Triangulo{Base: 12, Altura: 6}, areaEsperada: 36.0},
		}
	
		for _, tt := range testesArea {
			resultado := tt.forma.Area()
			if resultado != tt.areaEsperada {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.areaEsperada)
			}
		}
	})
}
func TestPerimetro(t *testing.T)  {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}
