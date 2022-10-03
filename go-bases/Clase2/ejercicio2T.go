package main

import "fmt"

/*
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
	Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
	Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del
ancho, si es cuadrática y cuál es el valor máximo.
*/

type Matrix struct {
	valores    [][]float64
	alto       int
	ancho      int
	cuadratica bool
	maximo     float64
}

func (m Matrix) Set() {
	if m.alto == m.ancho {
		m.cuadratica = true
	}
}

func (m Matrix) Print() {
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			fmt.Print("\t", m.valores[i][j])
		}
		fmt.Println("")
	}
}

func (m Matrix) Maximo() float64 {
	var maximo float64

	maximo = m.valores[0][0]

	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			if maximo < m.valores[i][j] {
				maximo = m.valores[i][j]
			}
		}
	}

	return maximo
}

func main() {

	matrizEjemplo := Matrix{
		valores: [][]float64{{1, 2, 3}, {4, 94, 65}, {76, 87, 87}},
		alto:    3,
		ancho:   3,
	}

	matrizEjemplo.Set()
	matrizEjemplo.Print()
	fmt.Println("Maximo: ", matrizEjemplo.Maximo())
	fmt.Println("Cuadratica?: ", matrizEjemplo.cuadratica)

}
