package goarea

import "math"

const pi = 3.1416

// Circ é responsável por calcular a área da circunferência.
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// TrianguloEq é responsável por calcular a área de um triângulo equilátero.
func TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
