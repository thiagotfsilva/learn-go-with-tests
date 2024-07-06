package estrutura

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Lagura float64
	Altura float64
}

// func (nomeDoReceptor TipoDoReceptor) nomeDoMetodo(argumentos)
func (r Retangulo) Area() float64 {
	return r.Lagura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Altura * t.Base) * 0.5
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Lagura)
}

/* func Area(retangulo Retangulo) float64 {
	return retangulo.Lagura * retangulo.Altura
} */